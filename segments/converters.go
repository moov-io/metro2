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

// converters handles golang to ACH type Converters
type converters struct{}

func (c *converters) parseValue(field Field, data string) interface{} {
	if field.Type&Numeric > 0 {
		value64, err := strconv.ParseInt(data, 10, 64)
		if err != nil {
			return nil
		}
		return value64
	} else if field.Type&Alphanumeric > 0 {
		var value = data
		return value
	} else if field.Type&Alpha > 0 {
		var value = strings.ToUpper(data)
		return value
	}

	return nil
}

func (c *converters) fillString(field Field) string {
	if field.Type&ZeroFill > 0 {
		return strings.Repeat("0", field.Length)
	}
	return strings.Repeat(" ", field.Length)
}

func (c *converters) toString(field Field, data reflect.Value) string {
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

func (c *converters) toSpecifications(fieldsFormat map[string]Field) []Specification {
	var specifications []Specification
	for key, field := range fieldsFormat {
		specifications = append(specifications, Specification{field.Start, key, field})
	}
	sort.Slice(specifications, func(i, j int) bool {
		return specifications[i].Key < specifications[j].Key
	})
	return specifications
}
