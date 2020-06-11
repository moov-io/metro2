// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/moov-io/ach"
)

var (
	upperAlphanumericRegex = regexp.MustCompile(`[^ A-Z0-9!"#$%&'()*+,-.\\/:;<>=?@\[\]^_{}|~]+`)
	alphanumericRegex      = regexp.MustCompile(`[^ \w!"#$%&'()*+,-.\\/:;<>=?@\[\]^_{}|~]+`)
	phoneRegex             = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	numericRegex           = regexp.MustCompile(`[0-9a-fA-F]`)
)

type validator struct{}

func (v *validator) isUpperalphanumeric(s string) error {
	if upperAlphanumericRegex.MatchString(s) {
		return ach.ErrUpperAlpha
	}
	return nil
}

func (v *validator) isAlphanumeric(s string) error {
	if alphanumericRegex.MatchString(s) {
		return ach.ErrNonAlphanumeric
	}
	return nil
}

func (v *validator) isNumeric(s string) error {
	if !numericRegex.MatchString(s) {
		return ErrNumeric
	}
	return nil
}

func (v *validator) isPhoneNumber(number int64) error {
	phoneNumber := fmt.Sprintf("%010d", number)
	if !phoneRegex.MatchString(phoneNumber) {
		return ErrPhoneNumber
	}
	return nil
}

func (v *validator) isValidType(elm field, data string) error {
	if elm.Required == required {
		if elm.Type&numeric > 0 {
			val, _ := strconv.Atoi(data)
			if val == 0 {
				return ach.ErrFieldRequired
			}
		} else if elm.Type&alphanumeric > 0 || elm.Type&alpha > 0 || elm.Type&descriptor > 0 {
			if len(data) == 0 {
				return ach.ErrFieldRequired
			}
		}
	}

	if elm.Type&numeric > 0 {
		return v.isNumeric(data)
	} else if elm.Type&alphanumeric > 0 {
		return v.isAlphanumeric(data)
	} else if elm.Type&alpha > 0 {
		return v.isUpperalphanumeric(data)
	} else if elm.Type&descriptor > 0 || elm.Type&packedDate > 0 || elm.Type&packedNumber > 0 ||
		elm.Type&packedTimestamp > 0 || elm.Type&timestamp > 0 || elm.Type&date > 0 {
		return nil
	}

	return ErrValidField
}

func (v *validator) validateFuncName(name string) string {
	return "Validate" + name
}

func newErrValidValue(field string) error {
	return fmt.Errorf("is an invalid value of %s", field)
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
