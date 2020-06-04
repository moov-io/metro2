// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"bufio"
	"gopkg.in/check.v1"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

type S struct {
	sampleBaseSegment string
}

var _ = check.Suite(&S{})

func (s *S) SetUpSuite(c *check.C) {
	f, err := os.Open(filepath.Join("../testdata", "base_segment.dat"))
	c.Assert(err, check.IsNil)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	s.sampleBaseSegment = strings.Join(lines, "")
}

func (s *S) TearDownSuite(c *check.C) {}

func (s *S) SetUpTest(c *check.C) {}

func (s *S) TearDownTest(c *check.C) {}

func (s *S) TestBaseSegment(c *check.C) {
	base := BaseSegment{}
	err := base.Parse(s.sampleBaseSegment)
	c.Assert(err, check.IsNil)
	err = base.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(base.String(), check.Equals, s.sampleBaseSegment)
}
