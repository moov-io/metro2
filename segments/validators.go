// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	upperalphanumericRegex = regexp.MustCompile(`[^ A-Z0-9!"#$%&'()*+,-.\\/:;<>=?@\[\]^_{}|~]+`)
	alphanumericRegex      = regexp.MustCompile(`[^ \w!"#$%&'()*+,-.\\/:;<>=?@\[\]^_{}|~]+`)
	phoneRegex             = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	numericRegex           = regexp.MustCompile(`[0-9a-fA-F]`)
	timestampRegex         = regexp.MustCompile(`(0[1-9]|1[0-2])(0[1-9]|[1-2][0-9]|3[0-1])[0-9]{4}(2[0-3]|[01][0-9])[0-5][0-9][0-5][0-9]`) //  MMDDYYYYHHMMSS
	dateRegex              = regexp.MustCompile(`(0[1-9]|1[0-2])(0[1-9]|[1-2][0-9]|3[0-1])[0-9]{4}`)                                       //  MMDDYYYY
)

type validator struct{}

func (v *validator) isUpperalphanumeric(s string) error {
	if upperalphanumericRegex.MatchString(s) {
		return ErrUpperAlpha
	}
	return nil
}

func (v *validator) isAlphanumeric(s string) error {
	if alphanumericRegex.MatchString(s) {
		return ErrAlphanumeric
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

func (v *validator) filledString(s string) (string, error) {
	if strings.Count(s, zeroString) == len(s) {
		return zeroString, nil
	}
	if strings.Count(s, nineString) == len(s) {
		return nineString, nil
	}
	if strings.Count(s, blankString) == len(s) {
		return blankString, nil
	}
	return "", errors.New("not filled")
}

func (v *validator) isValidType(elm field, data string) error {
	if elm.Required == required {
		if elm.Type&numeric > 0 {
			val, _ := strconv.Atoi(data)
			if val == 0 {
				return ErrRequired
			}
		} else if elm.Type&alphanumeric > 0 || elm.Type&alpha > 0 || elm.Type&binaryDescriptor > 0 {
			if len(data) == 0 {
				return ErrRequired
			}
		}
	}

	if elm.Type&numeric > 0 {
		return v.isNumeric(data)
	} else if elm.Type&alphanumeric > 0 {
		return v.isAlphanumeric(data)
	} else if elm.Type&alpha > 0 {
		return v.isUpperalphanumeric(data)
	} else if elm.Type&binaryDescriptor > 0 || elm.Type&packedDate > 0 ||
		elm.Type&packedNumber > 0 || elm.Type&packedDateLong > 0 {
		return nil
	}

	return ErrSegmentInvalidType
}

func (v *validator) isTimestamp(timestamp int64) error {
	timestampStr := fmt.Sprintf("%014d", timestamp)
	if !timestampRegex.MatchString(timestampStr) {
		return ErrTimestamp
	}
	return nil
}

func (v *validator) isDate(date int) error {
	dateStr := fmt.Sprintf("%08d", date)
	if !dateRegex.MatchString(dateStr) {
		return ErrDate
	}
	return nil
}

func (v *validator) validateFuncName(name string) string {
	return "Validate" + name
}
