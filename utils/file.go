package utils

import (
	"bufio"
	"os"
	"strconv"
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
	// check record identifier for header, trailer record
	if s[4] > 0x40 {
		return false
	}

	// packed format
	if s[6] == 0x00 && s[7] == 0x00 {
		return true
	}

	// unpacked format
	bdw, err := strconv.Atoi(s[0:4])
	if err != nil {
		return false
	}
	rdw, err := strconv.Atoi(s[4:8])
	if err != nil {
		return false
	}

	if rdw+4 == bdw {
		return true
	}

	return false
}

const (
	packedRecordLength = 366
	trailerIdentifier  = "TRAILER"
	headerIdentifier   = "HEADER"
)

// Metro file check
func IsMetroFile(s string) bool {
	if len(s) < packedRecordLength {
		return false
	}
	if s[4:10] == trailerIdentifier || s[8:14] == headerIdentifier {
		return true
	}
	return false
}
