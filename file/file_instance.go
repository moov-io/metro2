package file

import (
	"encoding/json"
	"strings"

	"github.com/moov-io/metro2/lib"
	"github.com/moov-io/metro2/utils"
)

// File contains the structures of a parsed metro 2 file.
type fileInstance struct {
	Header  lib.Record   `json:"header"`
	Bases   []lib.Record `json:"data"`
	Trailer lib.Record   `json:"trailer"`

	format string
}

// SetRecord can set block record like as header, trailer
func (f *fileInstance) SetRecord(r lib.Record) error {
	err := r.Validate()
	if err != nil {
		return err
	}

	switch r.Name() {
	case lib.HeaderRecordName, lib.PackedHeaderRecordName:
		f.Header = r
	case lib.TrailerRecordName, lib.PackedTrailerRecordName:
		f.Trailer = r
	default:
		return utils.NewErrValidRecord(r.Name())
	}

	return nil
}

// AddDataRecord can append data record
func (f *fileInstance) AddDataRecord(r lib.Record) error {
	err := r.Validate()
	if err != nil {
		return err
	}

	switch r.Name() {
	case lib.BaseSegmentName, lib.PackedBaseSegmentName:
		f.Bases = append(f.Bases, r)
	default:
		return utils.NewErrValidRecord(r.Name())
	}

	return nil
}

// GetRecord returns single record like as header, trailer.
func (f *fileInstance) GetRecord(name string) (lib.Record, error) {
	switch name {
	case HeaderRecordName:
		return f.Header, nil
	case TrailerRecordName:
		return f.Trailer, nil
	default:
		return nil, utils.NewErrValidRecord(name)
	}
}

// GetDataRecords returns data records
func (f *fileInstance) GetDataRecords() []lib.Record {
	return f.Bases
}

// GeneratorTrailer returns trailer segment that created automatically
func (f *fileInstance) GeneratorTrailer() (lib.Record, error) {
	return nil, nil
}

// Validate performs some checks on the file and returns an error if not Validated
func (f *fileInstance) Validate() error {
	return nil
}

// Parse attempts to initialize a *File object assuming the input is valid raw data.
func (f *fileInstance) Parse(record string) error {
	f.Bases = []lib.Record{}
	offset := 0

	// Header Record
	hread, err := f.Header.Parse(record)
	if err != nil {
		return err
	}
	offset += hread

	// Data Record
	for err == nil {
		var base lib.Record
		if f.format == PackedFileFormat {
			base = lib.NewPackedBaseSegment()
		} else {
			base = lib.NewBaseSegment()
		}

		read, err := base.Parse(record[offset:])
		if err != nil {
			break
		}
		f.Bases = append(f.Bases, base)
		offset += read
	}

	// Trailer Record
	tread, err := f.Trailer.Parse(record[offset:])
	if err != nil {
		return err
	}
	offset += tread

	if offset != len(record) {
		return utils.NewErrParse()
	}

	return nil
}

// String writes the File struct to raw string.
func (f *fileInstance) String() string {
	var buf strings.Builder

	// Header Block
	header := f.Header.String()

	// Data Block
	data := ""
	for _, base := range f.Bases {
		data += base.String()
	}

	// Trailer Block
	trailer := f.Trailer.String()

	buf.Grow(len(header) + len(data) + len(trailer))
	buf.WriteString(header)
	buf.WriteString(data)
	buf.WriteString(trailer)

	return buf.String()
}

// UnmarshalJSON parses a JSON blob
func (f *fileInstance) UnmarshalJSON(data []byte) error {

	dummy := make(map[string]interface{})
	err := json.Unmarshal(data, &dummy)
	if err != nil {
		return err
	}

	for name, record := range dummy {
		buf, err := json.Marshal(record)
		if err != nil {
			return err
		}

		switch name {
		case HeaderRecordName:
			err = json.Unmarshal(buf, f.Header)
			if err != nil {
				return err
			}
		case TrailerRecordName:
			err = json.Unmarshal(buf, f.Trailer)
			if err != nil {
				return err
			}
		case DataRecordName:
			var list []interface{}
			err = json.Unmarshal(buf, &list)
			if err != nil {
				return nil
			}
			for _, subSegment := range list {
				subBuf, err := json.Marshal(subSegment)
				if err != nil {
					return err
				}
				if f.format == CharacterFileFormat {
					base := lib.NewBaseSegment()
					err = json.Unmarshal(subBuf, base)
					if err != nil {
						return nil
					}
					f.Bases = append(f.Bases, base)
				} else {
					base := lib.NewPackedBaseSegment()
					err = json.Unmarshal(subBuf, base)
					if err != nil {
						return nil
					}
					f.Bases = append(f.Bases, base)
				}
			}
		}
	}

	return nil
}
