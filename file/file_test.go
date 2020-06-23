package file

import (
	"bytes"
	"encoding/json"
	"github.com/moov-io/metro2/lib"
	"gopkg.in/check.v1"
	"strings"
)

func (t *FileTest) TestJsonWithUnpackedVariableBlocked(c *check.C) {
	f, err := NewFile(CharacterFileFormat)
	c.Assert(err, check.IsNil)
	err = json.Unmarshal([]byte(t.unpackedVariableBlockedJson), f)
	c.Assert(err, check.IsNil)
	c.Assert(f.String(), check.Equals, t.unpackedVariableBlockedFile)
	buf, err := json.Marshal(f)
	c.Assert(err, check.IsNil)
	var out bytes.Buffer
	err = json.Indent(&out, buf, "", "  ")
	c.Assert(err, check.IsNil)
	jsonStr := out.String()
	jsonStr = strings.ReplaceAll(jsonStr, "\n", "")
	c.Assert(jsonStr, check.Equals, t.unpackedVariableBlockedJson)
}

func (t *FileTest) TestParseWithUnpackedVariableBlockedFileParse(c *check.C) {
	f, err := NewFile(CharacterFileFormat)
	c.Assert(err, check.IsNil)
	err = f.Parse(t.unpackedVariableBlockedFile)
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestJsonWithUnpackedFixedLength(c *check.C) {
	f, err := NewFile(CharacterFileFormat)
	c.Assert(err, check.IsNil)
	err = json.Unmarshal([]byte(t.unpackedFixedLengthJson), f)
	c.Assert(err, check.IsNil)
	c.Assert(f.String(), check.Equals, t.unpackedFixedLengthFile)
	buf, err := json.Marshal(f)
	c.Assert(err, check.IsNil)
	var out bytes.Buffer
	err = json.Indent(&out, buf, "", "  ")
	c.Assert(err, check.IsNil)
	jsonStr := out.String()
	jsonStr = strings.ReplaceAll(jsonStr, "\n", "")
	c.Assert(jsonStr, check.Equals, t.unpackedFixedLengthJson)
}

func (t *FileTest) TestParseWithUnpackedFixedLength(c *check.C) {
	f, err := NewFile(CharacterFileFormat)
	c.Assert(err, check.IsNil)
	err = f.Parse(t.unpackedFixedLengthFile)
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestJsonWithPackedBlocked(c *check.C) {
	f, err := NewFile(PackedFileFormat)
	c.Assert(err, check.IsNil)
	err = json.Unmarshal([]byte(t.packedJson), f)
	c.Assert(err, check.IsNil)
	c.Assert(f.String(), check.Equals, t.packedFile)
	buf, err := json.Marshal(f)
	c.Assert(err, check.IsNil)
	var out bytes.Buffer
	err = json.Indent(&out, buf, "", "  ")
	c.Assert(err, check.IsNil)
	jsonStr := out.String()
	jsonStr = strings.ReplaceAll(jsonStr, "\n", "")
	c.Assert(jsonStr, check.Equals, t.packedJson)
}

func (t *FileTest) TestParseWithPackedFileParse(c *check.C) {
	f, err := NewFile(PackedFileFormat)
	c.Assert(err, check.IsNil)
	err = f.Parse(t.packedFile)
	c.Assert(err, check.IsNil)
}

func (t *FileTest) TestFileSetBlock(c *check.C) {
	f, err := NewFile(CharacterFileFormat)
	c.Assert(err, check.IsNil)
	jsonStr := `{
		"blockDescriptorWord": 430,
		"recordDescriptorWord": 426,
		"recordIdentifier": "HEADER",
		"transUnionProgramIdentifier": "5555555555",
		"activityDate": "2002-08-20T00:00:00Z",
		"dateCreated": "1999-05-10T00:00:00Z",
		"programDate": "1999-05-10T00:00:00Z",
		"programRevisionDate": "1999-05-10T00:00:00Z",
		"reporterName": "YOUR BUSINESS NAME HERE",
		"reporterAddress": "LINE ONE OF YOUR ADDRESS LINE TWO OF YOUR ADDRESS LINE THERE OF YOUR ADDRESS",
		"reporterTelephoneNumber": 1234567890
	  }`
	newSegment := lib.HeaderRecord{}
	err = json.Unmarshal([]byte(jsonStr), &newSegment)
	c.Assert(err, check.IsNil)
	orgHeader, err := f.GetRecord(HeaderRecordName)
	c.Assert(err, check.IsNil)
	origin := orgHeader.BlockSize()
	err = f.SetRecord(&newSegment)
	c.Assert(err, check.IsNil)
	newHeader, err := f.GetRecord(HeaderRecordName)
	c.Assert(err, check.IsNil)
	c.Assert(origin, check.Not(check.Equals), newHeader.BlockSize())
}

func (t *FileTest) TestFileDataRecord(c *check.C) {
	f, err := NewFile(CharacterFileFormat)
	c.Assert(err, check.IsNil)
	segment := lib.NewBaseSegment()
	err = json.Unmarshal([]byte(t.baseSegmentJson), &segment)
	c.Assert(err, check.IsNil)
	err = f.AddDataRecord(segment)
	c.Assert(err, check.IsNil)
	list := f.GetDataRecords()
	c.Assert(len(list), check.Equals, 1)
}
