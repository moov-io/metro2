// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"testing"
)

func TestValidatorUpperAlpha(t *testing.T) {
	v := &validator{}
	if err := v.isUpperAlphanumeric("ab91", "", ""); err == nil {
		t.Error("expected error")
	}
	if err := v.isUpperAlphanumeric("AB91", "", ""); err != nil {
		t.Error(err)
	}
}

func TestValidatorFilledString(t *testing.T) {
	if !validFilledString("9") {
		t.Error("expected error")
	}
	if !validFilledString(" ") {
		t.Error("expected error")
	}
}

func TestValidatorAlphanumeric(t *testing.T) {
	v := &validator{}
	if err := v.isAlphanumeric("Ϡϛβ123", "", ""); err == nil {
		t.Error("expected error")
	}
}

func TestIsValidType(t *testing.T) {
	v := &validator{}
	test1 := field{
		Required: required,
		Type:     alphanumeric,
		Length:   9,
	}
	if err := v.isValidType(test1, "", "", ""); err == nil {
		t.Error("expected error")
	}
	test1.Type = date
	if err := v.isValidType(test1, "999999999", "", ""); err == nil {
		t.Error("expected error")
	}
	test1.Type = 1 << 15
	if err := v.isValidType(test1, "999999999", "", ""); err == nil {
		t.Error("expected error")
	}
}
