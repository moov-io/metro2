package utils

import (
	"bufio"
	"os"
	"strings"
)

func ReadFile(f *os.File) string {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return strings.Join(lines, "")
}
