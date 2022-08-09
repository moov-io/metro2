// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/moov-io/metro2/pkg/utils"
)

type converter struct{}

func (c *converter) parseValue(elm field, data, fieldName, recordName string) (reflect.Value, error) {
	if elm.Type&numeric > 0 {
		value, err := strconv.ParseInt(data, 10, 64)
		return reflect.ValueOf(value), err
	} else if elm.Type&timestamp > 0 {
		ret, err := timeFromTimestampString(data)
		return reflect.ValueOf(ret), err
	} else if elm.Type&date > 0 {
		ret, err := timeFromDateString(data)
		return reflect.ValueOf(ret), err
	} else if elm.Type&alphanumeric > 0 {
		data = strings.TrimRight(data, blankString)
		return reflect.ValueOf(data), nil
	} else if elm.Type&alpha > 0 {
		return reflect.ValueOf(strings.ToUpper(data)), nil
	} else if elm.Type&descriptor > 0 {
		value := int64(binary.BigEndian.Uint16([]byte(data)))
		return reflect.ValueOf(value), nil
	} else if elm.Type&packedTimestamp > 0 {
		ret, err := timeFromPackedTimestampString(data)
		return reflect.ValueOf(ret), err
	} else if elm.Type&packedDate > 0 {
		ret, err := timeFromPackedDateString(data)
		return reflect.ValueOf(ret), err
	} else if elm.Type&packedNumber > 0 {
		return reflect.ValueOf(packedNumberFromString(data)), nil
	}

	return reflect.Value{}, utils.NewErrInvalidValueOfField(fieldName, recordName)
}

func (c *converter) fillString(elm field) string {
	if elm.Type&zeroFill > 0 {
		return strings.Repeat(zeroString, elm.Length)
	}
	return strings.Repeat(blankString, elm.Length)
}

func (c *converter) toString(elm field, data reflect.Value) string {
	if !data.IsValid() {
		return c.fillString(elm)
	}
	if elm.Type&omitted > 0 && data.Interface().(int) == 0 {
		return ""
	}

	sizeStr := strconv.Itoa(elm.Length)
	if elm.Type&numeric > 0 {
		return fmt.Sprintf("%0"+sizeStr+"."+sizeStr+"v", fmt.Sprintf("%v", data))
	} else if elm.Type&timestamp > 0 {
		if datatime, ok := data.Interface().(utils.Time); ok {
			if t := time.Time(datatime); !t.IsZero() {
				return t.Format(timestampFormat)
			}
		}
		return strings.Repeat(zeroString, elm.Length)
	} else if elm.Type&date > 0 {
		if datatime, ok := data.Interface().(utils.Time); ok {
			if t := time.Time(datatime); !t.IsZero() {
				return t.Format(dateFormat)
			}
		}
		return strings.Repeat(zeroString, elm.Length)
	} else if elm.Type&alphanumeric > 0 || elm.Type&alpha > 0 {
		return fmt.Sprintf("%-"+sizeStr+"."+sizeStr+"s", data)
	} else if elm.Type&descriptor > 0 {
		return descriptorString(data)
	} else if elm.Type&packedTimestamp > 0 {
		return packedTimeString(data, timestampFormat, elm.Length, packedTimestampSize)
	} else if elm.Type&packedDate > 0 {
		return packedTimeString(data, dateFormat, elm.Length, packedDateSize)
	} else if elm.Type&packedNumber > 0 {
		return packedNumberString(data, elm.Length)
	}

	return c.fillString(elm)
}

func (c *converter) toSpecifications(fieldsFormat map[string]field) []specification {
	var specifications []specification
	for key, field := range fieldsFormat {
		specifications = append(specifications, specification{field.Start, key, field})
	}
	sort.Slice(specifications, func(i, j int) bool {
		if specifications[i].Key == specifications[j].Key {
			return specifications[i].Name < specifications[j].Name
		}
		return specifications[i].Key < specifications[j].Key
	})
	return specifications
}

// parse field with string
func (c *converter) parseRecordValues(fields reflect.Value, spec map[string]field, record []byte, v *validator, recordName string) (int, error) {
	offset := 0
	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		// skip local variable
		if !unicode.IsUpper([]rune(fieldName)[0]) {
			continue
		}
		field := fields.FieldByName(fieldName)
		spec, ok := spec[fieldName]
		if !ok || !field.IsValid() {
			return 0, utils.NewErrInvalidValueOfField(fieldName, recordName)
		}

		if len(record) < spec.Start+spec.Length+offset {
			return 0, utils.NewErrSegmentLength(recordName)
		}
		data := string(record[spec.Start+offset : spec.Start+spec.Length+offset])
		if err := v.isValidType(spec, data, fieldName, recordName); err != nil {
			return 0, err
		}

		value, err := c.parseValue(spec, data, fieldName, recordName)
		if err != nil {
			return 0, err
		}
		// set value
		if value.IsValid() && field.CanSet() {
			switch value.Interface().(type) {
			case int, int64:
				if fieldName == "BlockDescriptorWord" {
					if !utils.IsVariableLength(record) {
						continue
					}
					offset += 4
				}
				field.SetInt(value.Interface().(int64))
			case string:
				field.SetString(value.Interface().(string))
			case utils.Time:
				field.Set(value)
			}
		}
	}
	return 0, nil
}

// convert functions
func timeFromTimestampString(date string) (utils.Time, error) {
	if strings.Count(date, "0") != len(date) {
		time, err := time.Parse(timestampFormat, date)
		return utils.Time(time), err
	}
	return utils.Time{}, nil
}

func timeFromDateString(date string) (utils.Time, error) {
	if strings.Count(date, "0") != len(date) {
		time, err := time.Parse(dateFormat, date)
		return utils.Time(time), err
	}
	return utils.Time{}, nil
}

func timeFromPackedTimestampString(date string) (utils.Time, error) {
	value := int64(0)
	bin := []byte(date)
	if bin[0] == 0x00 && bin[packedTimestampSize-1] == 0x73 {
		var in bytes.Buffer
		in.Grow(int64size)
		for i := 0; i < 2; i++ {
			in.WriteByte(0x00)
		}
		in.Write(bin[1 : packedTimestampSize-1])
		value = int64(binary.BigEndian.Uint64(in.Bytes()))
	}

	datestr := fmt.Sprintf("%0"+timestampSizeStr+"d", value)
	return timeFromTimestampString(datestr)
}

func timeFromPackedDateString(date string) (utils.Time, error) {
	value := int64(0)
	bin := []byte(date)
	if bin[0] == 0x00 && bin[packedDateSize-1] == 0x73 {
		var in bytes.Buffer
		in.Grow(int64size)
		for i := 0; i < 5; i++ {
			in.WriteByte(0x00)
		}
		in.Write(bin[1 : packedDateSize-1])
		value = int64(binary.BigEndian.Uint64(in.Bytes()))
	}

	datestr := fmt.Sprintf("%0"+dateSizeStr+"d", value)
	return timeFromDateString(datestr)
}

func packedNumberFromString(data string) int64 {
	length := len(data)
	var in bytes.Buffer

	in.Grow(int64size)
	for i := 0; i < int64size-length; i++ {
		in.WriteByte(0x00)
	}
	in.Write([]byte(data))
	value := int64(binary.BigEndian.Uint64(in.Bytes()))
	return value
}

func packedTimeString(data reflect.Value, format string, length int, size int) string {
	value := int64(0)
	if data.Type() == reflect.TypeOf(utils.Time{}) {
		if data.Interface() == nil {
			return ""
		}

		newTime := data.Interface().(utils.Time)
		value, _ = strconv.ParseInt(time.Time(newTime).Format(format), 10, 64)
	}

	var out bytes.Buffer
	out.Grow(length)
	if value > 0 {
		out.WriteByte(0x00)
		v := uint64(value)
		for i := 0; i < size-2; i++ {
			out.WriteByte(byte(v >> (8 * (size - 3 - i))))
		}
		out.WriteByte(0x73)
	} else {
		for i := 0; i < size; i++ {
			out.WriteByte(0x00)
		}
	}
	return out.String()
}

func packedNumberString(data reflect.Value, length int) string {
	var out bytes.Buffer
	out.Grow(length)
	if data.Int() > 0 {
		v := uint64(data.Int())
		for i := 0; i < length; i++ {
			shift := 8 * (length - i - 1)
			if shift > 0 {
				out.WriteByte(byte(v >> shift))
			} else {
				out.WriteByte(byte(v))
			}
		}
	} else {
		for i := 0; i < length; i++ {
			out.WriteByte(0x00)
		}
	}
	return out.String()
}

func descriptorString(data reflect.Value) string {
	value := make([]byte, 4)
	binary.BigEndian.PutUint16(value[0:], uint16(data.Int()))
	return string(value)
}
