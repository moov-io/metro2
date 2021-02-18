// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/moov-io/metro2/pkg/utils"
)

type validator struct{}

func (v *validator) isUpperAlphanumeric(s, fieldName, recordName string) error {
	if upperAlphanumericRegex.MatchString(s) {
		return utils.NewErrUpperAlpha(fieldName, recordName)
	}
	return nil
}

func (v *validator) isAlphanumeric(s, filedName, recordName string) error {
	if alphanumericRegex.MatchString(s) {
		return utils.NewErrNonAlphanumeric(filedName, recordName)
	}
	return nil
}

func (v *validator) isNumeric(s, filedName, recordName string) error {
	if !numericRegex.MatchString(s) {
		return utils.NewErrNumeric(filedName, recordName)
	}
	return nil
}

func (v *validator) isPhoneNumber(number int64, recordName string) error {
	phoneNumber := fmt.Sprintf("%010d", number)
	if !phoneRegex.MatchString(phoneNumber) {
		return utils.NewErrPhoneNumber(recordName)
	}
	return nil
}

func (v *validator) isValidType(elm field, data, fieldName, recordName string) error {
	// required check
	if elm.Required == required {
		if elm.Type&numeric > 0 {
			val, _ := strconv.Atoi(data)
			if val == 0 {
				return utils.NewErrFieldRequired(fieldName, recordName)
			}
		} else if elm.Type&alphanumeric > 0 || elm.Type&alpha > 0 || elm.Type&descriptor > 0 {
			if len(data) == 0 {
				return utils.NewErrFieldRequired(fieldName, recordName)
			}
		} else if elm.Type&timestamp > 0 || elm.Type&date > 0 {
			if validFilledString(data) && len(data) == elm.Length {
				return utils.NewErrFieldRequired(fieldName, recordName)
			}
		}
	}

	// date check
	if elm.Type&numeric > 0 {
		return v.isNumeric(data, fieldName, recordName)
	} else if elm.Type&alphanumeric > 0 {
		return v.isAlphanumeric(data, fieldName, recordName)
	} else if elm.Type&alpha > 0 {
		return v.isUpperAlphanumeric(data, fieldName, recordName)
	} else if elm.Type&descriptor > 0 || elm.Type&packedDate > 0 || elm.Type&packedNumber > 0 ||
		elm.Type&packedTimestamp > 0 || elm.Type&timestamp > 0 || elm.Type&date > 0 {
		return nil
	}

	return utils.NewErrInvalidValueOfField(fieldName, recordName)
}

func (v *validator) validateFuncName(name string) string {
	return "Validate" + name
}

// to validate fields of record
func (v *validator) validateRecord(r interface{}, spec map[string]field, recordName string) error {
	fields := reflect.ValueOf(r).Elem()
	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		if spec, ok := spec[fieldName]; ok {
			if spec.Required == required {
				fieldValue := fields.FieldByName(fieldName)
				if fieldValue.IsZero() {
					return utils.NewErrFieldRequired(fieldName, recordName)
				}
			}
		}

		funcName := v.validateFuncName(fieldName)
		method := reflect.ValueOf(r).MethodByName(funcName)
		if method.IsValid() {
			response := method.Call(nil)
			if len(response) == 0 {
				continue
			}

			err := method.Call(nil)[0]
			if !err.IsNil() {
				return err.Interface().(error)
			}
		}
	}

	return nil
}

func validFilledString(s string) bool {
	if strings.Count(s, zeroString) == len(s) {
		return true
	}
	if strings.Count(s, nineString) == len(s) {
		return true
	}
	if strings.Count(s, blankString) == len(s) {
		return true
	}
	return false
}
