// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/moov-io/metro2/segments"
	"gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

type FileTest struct {
}

var _ = check.Suite(&FileTest{})

func (s *FileTest) SetUpSuite(c *check.C) {
}

func (s *FileTest) TearDownSuite(c *check.C) {}

func (s *FileTest) SetUpTest(c *check.C) {}

func (s *FileTest) TearDownTest(c *check.C) {}

func (s *FileTest) TestFileUnmarshal(c *check.C) {}

func (s *FileTest) TestFileMarshal(c *check.C) {
	f, err := NewFile(CharacterFileFormat)
	c.Assert(err, check.IsNil)
	f.AddApplicableSegment(segments.NewN1Segment())
	f.AddApplicableSegment(segments.NewL1Segment())
	jsonStr, err := json.Marshal(f)
	c.Assert(err, check.IsNil)
	var out bytes.Buffer
	err = json.Indent(&out, jsonStr, "", "\t")
	c.Assert(err, check.IsNil)
	fmt.Println(out.String())
}
