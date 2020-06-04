// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

const (
	zeroString  = "0"
	blankString = " "
	nineString  = "9"
)

type converter struct{}

func (c *converter) parseValue(field Field, data string) (reflect.Value, error) {
	if field.Type&Numeric > 0 {
		value, err := strconv.ParseInt(data, 10, 64)
		if err != nil {
			return reflect.Value{}, err
		}
		return  reflect.ValueOf(value), nil
	} else if field.Type&Alphanumeric > 0 {
		return  reflect.ValueOf(data), nil
	} else if field.Type&Alpha > 0 {
		upperString := strings.ToUpper(data)
		return  reflect.ValueOf(upperString), nil
	}

	return reflect.Value{}, ErrSegmentParseType
}

func (c *converter) fillString(field Field) string {
	if field.Type&ZeroFill > 0 {
		return strings.Repeat(zeroString, field.Length)
	}
	return strings.Repeat(blankString, field.Length)
}

func (c *converter) toString(field Field, data reflect.Value) string {
	if !data.IsValid() {
		return c.fillString(field)
	}
	if field.Type&Omitted > 0 {
		return ""
	}

	fieldSize := strconv.Itoa(field.Length)
	if field.Type&Numeric > 0 {
		return fmt.Sprintf("%0"+fieldSize+"d", data)
	} else if field.Type&Alphanumeric > 0 || field.Type&Alpha > 0 {
		return fmt.Sprintf("%-"+fieldSize+"s", data)
	}

	return c.fillString(field)
}

func (c *converter) toSpecifications(fieldsFormat map[string]Field) []Specification {
	var specifications []Specification
	for key, field := range fieldsFormat {
		specifications = append(specifications, Specification{field.Start, key, field})
	}
	sort.Slice(specifications, func(i, j int) bool {
		return specifications[i].Key < specifications[j].Key
	})
	return specifications
}
