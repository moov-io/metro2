// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/moov-io/metro2/utils"
)

type validator struct{}

func (v *validator) isUpperAlphanumeric(s string) error {
	if upperAlphanumericRegex.MatchString(s) {
		return utils.ErrUpperAlpha
	}
	return nil
}

func (v *validator) isAlphanumeric(s string) error {
	if alphanumericRegex.MatchString(s) {
		return utils.ErrNonAlphanumeric
	}
	return nil
}

func (v *validator) isNumeric(s string) error {
	if !numericRegex.MatchString(s) {
		return utils.ErrNumeric
	}
	return nil
}

func (v *validator) isPhoneNumber(number int64) error {
	phoneNumber := fmt.Sprintf("%010d", number)
	if !phoneRegex.MatchString(phoneNumber) {
		return utils.ErrPhoneNumber
	}
	return nil
}

func (v *validator) isValidType(elm field, data string) error {
	// required check
	if elm.Required == required {
		if elm.Type&numeric > 0 {
			val, _ := strconv.Atoi(data)
			if val == 0 {
				return utils.ErrFieldRequired
			}
		} else if elm.Type&alphanumeric > 0 || elm.Type&alpha > 0 || elm.Type&descriptor > 0 {
			if len(data) == 0 {
				return utils.ErrFieldRequired
			}
		} else if elm.Type&timestamp > 0 || elm.Type&date > 0 {
			if validFilledString(data) {
				return utils.ErrFieldRequired
			}
		}
	}

	// date check
	if elm.Type&numeric > 0 {
		return v.isNumeric(data)
	} else if elm.Type&alphanumeric > 0 {
		return v.isAlphanumeric(data)
	} else if elm.Type&alpha > 0 {
		return v.isUpperAlphanumeric(data)
	} else if elm.Type&descriptor > 0 || elm.Type&packedDate > 0 || elm.Type&packedNumber > 0 ||
		elm.Type&packedTimestamp > 0 || elm.Type&timestamp > 0 || elm.Type&date > 0 {
		return nil
	}

	return utils.ErrValidField
}

func (v *validator) validateFuncName(name string) string {
	return "Validate" + name
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
