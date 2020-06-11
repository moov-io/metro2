// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	zeroString      = "0"
	blankString     = " "
	nineString      = "9"
	timestampFormat = "01022006150405"
	dateFormat      = "01022006"
)

type converter struct{}

func (c *converter) parseValue(elm field, data string) (reflect.Value, error) {
	if elm.Type&numeric > 0 {
		value, err := strconv.ParseInt(data, 10, 64)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(value), nil
	} else if elm.Type&timestamp > 0 {
		if !validFilledString(data) {
			value, err := time.Parse(timestampFormat, data)
			return reflect.ValueOf(value), err
		}
		return reflect.ValueOf(time.Time{}), nil
	} else if elm.Type&date > 0 {
		if !validFilledString(data) {
			value, err := time.Parse(dateFormat, data)
			return reflect.ValueOf(value), err
		}
		return reflect.ValueOf(time.Time{}), nil
	} else if elm.Type&alphanumeric > 0 {
		return reflect.ValueOf(data), nil
	} else if elm.Type&alpha > 0 {
		upperString := strings.ToUpper(data)
		return reflect.ValueOf(upperString), nil
	} else if elm.Type&descriptor > 0 {
		value := int64(binary.BigEndian.Uint16([]byte(data)))
		return reflect.ValueOf(value), nil
	} else if elm.Type&packedTimestamp > 0 {
		value := int64(0)
		bin := []byte(data)
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
		if !validFilledString(datestr) {
			value, err := time.Parse(timestampFormat, datestr)
			return reflect.ValueOf(value), err
		}
		return reflect.ValueOf(time.Time{}), nil
	} else if elm.Type&packedDate > 0 {
		value := int64(0)
		bin := []byte(data)
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
		if !validFilledString(datestr) {
			value, err := time.Parse(dateFormat, datestr)
			return reflect.ValueOf(value), err
		}
		return reflect.ValueOf(time.Time{}), nil
	} else if elm.Type&packedNumber > 0 {
		length := len(data)
		var in bytes.Buffer

		in.Grow(int64size)
		for i := 0; i < int64size-length; i++ {
			in.WriteByte(0x00)
		}
		in.Write([]byte(data))
		value := int64(binary.BigEndian.Uint64(in.Bytes()))
		return reflect.ValueOf(value), nil
	}

	return reflect.Value{}, ErrValidField
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
	if elm.Type&omitted > 0 {
		return ""
	}

	fieldSize := strconv.Itoa(elm.Length)
	if elm.Type&numeric > 0 {
		return fmt.Sprintf("%0"+fieldSize+"d", data)
	} else if elm.Type&timestamp > 0 {
		if datatime, ok := data.Interface().(time.Time); ok {
			if !datatime.IsZero() {
				return datatime.Format(timestampFormat)
			}
		}
		return strings.Repeat(zeroString, elm.Length)
	} else if elm.Type&date > 0 {
		if datatime, ok := data.Interface().(time.Time); ok {
			if !datatime.IsZero() {
				return datatime.Format(dateFormat)
			}
		}
		return strings.Repeat(zeroString, elm.Length)
	} else if elm.Type&alphanumeric > 0 || elm.Type&alpha > 0 {
		return fmt.Sprintf("%-"+fieldSize+"s", data)
	} else if elm.Type&descriptor > 0 {
		value := make([]byte, 4)
		binary.BigEndian.PutUint16(value[0:], uint16(data.Int()))
		return string(value)
	} else if elm.Type&packedDate > 0 {
		datastr := strings.Repeat(zeroString, elm.Length)
		if datatime, ok := data.Interface().(time.Time); ok {
			if !datatime.IsZero() {
				datastr = datatime.Format(dateFormat)
			}
		}
		dataint, err := strconv.Atoi(datastr)
		if err != nil {
			return datastr
		}
		var out bytes.Buffer
		out.Grow(elm.Length)
		if dataint > 0 {
			out.WriteByte(0x00)
			v := uint64(dataint)
			for i := 0; i < packedDateSize-2; i++ {
				out.WriteByte(byte(v >> (8 * (packedDateSize - 3 - i))))
			}
			out.WriteByte(0x73)
		} else {
			for i := 0; i < packedDateSize; i++ {
				out.WriteByte(0x00)
			}
		}
		return out.String()
	} else if elm.Type&packedTimestamp > 0 {
		datastr := strings.Repeat(zeroString, elm.Length)
		if datatime, ok := data.Interface().(time.Time); ok {
			if !datatime.IsZero() {
				datastr = datatime.Format(timestampFormat)
			}
		}
		dataint, err := strconv.Atoi(datastr)
		if err != nil {
			return datastr
		}
		var out bytes.Buffer
		out.Grow(elm.Length)
		if dataint > 0 {
			out.WriteByte(0x00)
			v := uint64(dataint)
			for i := 0; i < packedTimestampSize-2; i++ {
				out.WriteByte(byte(v >> (8 * (packedTimestampSize - 3 - i))))
			}
			out.WriteByte(0x73)
		} else {
			for i := 0; i < packedTimestampSize; i++ {
				out.WriteByte(0x00)
			}
		}
		return out.String()
	} else if elm.Type&packedNumber > 0 {
		var out bytes.Buffer
		out.Grow(elm.Length)
		if data.Int() > 0 {
			length := elm.Length
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
			for i := 0; i < elm.Length; i++ {
				out.WriteByte(0x00)
			}
		}
		return out.String()
	}

	return c.fillString(elm)
}

func (c *converter) toSpecifications(fieldsFormat map[string]field) []specification {
	var specifications []specification
	for key, field := range fieldsFormat {
		specifications = append(specifications, specification{field.Start, key, field})
	}
	sort.Slice(specifications, func(i, j int) bool {
		return specifications[i].Key < specifications[j].Key
	})
	return specifications
}
