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
func ReadFile(f *os.File) string {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return strings.Join(lines, "")
}

// Variable block check
func IsVariableLength(s string) bool {

	// Checking header record identifier
	if len(s) > 15 && strings.ToUpper(s[8:14]) == "HEADER" {
		return true
	}

	// Checking base record field 4
	//  Field formerly used for Correction Indicator.
	if s[17] == 0x30 {
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

// Pass the record offset to check its trailer or not
func IsTrailerRecord(s string) bool {
	if len(s) < packedRecordLength {
		return false
	}
	if s[4:11] == Trailerdentifier || s[8:15] == Trailerdentifier {
		return true
	}
	return false
}
