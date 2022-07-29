// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package utils

import (
	"bufio"
	"os"
	"strings"
)

// File Read
func ReadFile(f *os.File) []byte {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var raw []byte
	for scanner.Scan() {
		raw = append(raw, scanner.Bytes()...)
	}

	return raw
}

// Variable block check
func IsVariableLength(data []byte) bool {

	// Checking header record identifier
	if len(data) > 15 && strings.ToUpper(string(data[8:14])) == "HEADER" {
		return true
	}

	// Checking base record field 4
	//  Field formerly used for Correction Indicator.
	if len(data) > 18 && data[17] == 0x30 {
		return true
	}

	return false
}

// IsPacked packed format check
func IsPacked(s string) bool {

	// fix packed format
	if s[2] == 0x00 && s[3] == 0x00 {
		return true
	}

	// variable packed format
	if s[6] == 0x00 && s[7] == 0x00 {
		return true
	}

	return false
}

// Metro file check
func IsMetroFile(s string) bool {
	if len(s) < packedRecordLength {
		return false
	}
	if s[4:10] == headerIdentifier || s[8:14] == headerIdentifier {
		return true
	}
	return false
}
