package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/moov-io/metro2/lib"
	"github.com/moov-io/metro2/utils"
	"os"
	"path/filepath"
)

func main() {
	f, err := os.Open(filepath.Join(".", "testdata", "base_segment.dat"))
	if err != nil {
		return
	}
	sample := utils.ReadFile(f)

	segment := lib.NewBaseSegment()
	_, err = segment.Parse(sample)
	if err != nil {
		return
	}
	err = segment.Validate()
	if err != nil {
		return
	}
	buf, err := json.Marshal(segment)
	if err != nil {
		return
	}

	var pretty bytes.Buffer
	err = json.Indent(&pretty, buf, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(pretty.String())
}
