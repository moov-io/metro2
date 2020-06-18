package file

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/moov-io/metro2/segments"
	"github.com/moov-io/metro2/utils"
)

// File contains the structures of a parsed metro 2 file.
type fileInstance struct {
	Header     segments.Segment            `json:"headerRecord"`
	Base       segments.Segment            `json:"baseSegment"`
	J1Segments []segments.Segment          `json:"j1Segments"`
	J2Segments []segments.Segment          `json:"jsSegments"`
	Appendages map[string]segments.Segment `json:"appendages"`
	Trailer    segments.Segment            `json:"trailerRecord"`

	format string
}

// SetBlock can set block record like as header, base, trailer
func (f *fileInstance) SetBlock(s segments.Segment) error {
	err := s.Validate()
	if err != nil {
		return err
	}
	switch s.Description() {
	case segments.HeaderRecordDescription:
		f.Header = s
	case segments.BaseSegmentDescription:
		f.Base = s
	case segments.TrailerRecordDescription:
		f.Trailer = s
	}
	return nil
}

// AddApplicableSegment can append applicable segment like as j1, j2, k1, k2, k3, k4, l1, n1
func (f *fileInstance) AddApplicableSegment(s segments.Segment) error {
	err := s.Validate()
	if err != nil {
		return err
	}
	switch s.Description() {
	case segments.J1SegmentDescription:
		f.J1Segments = append(f.J1Segments, s)
	case segments.J2SegmentDescription:
		f.J2Segments = append(f.J2Segments, s)
	case segments.K1SegmentDescription, segments.K2SegmentDescription, segments.K3SegmentDescription,
		segments.K4SegmentDescription, segments.L1SegmentDescription, segments.N1SegmentDescription:
		f.Appendages[s.Description()] = s
	}
	return nil
}

// GetSegment returns single segment like as header, base, trailer k1, k2, k3, k4, l1, n1
func (f *fileInstance) GetSegment(description string) segments.Segment {
	switch description {
	case segments.HeaderRecordDescription:
		return f.Header
	case segments.BaseSegmentDescription:
		return f.Base
	case segments.TrailerRecordDescription:
		return f.Trailer
	case segments.K1SegmentDescription, segments.K2SegmentDescription, segments.K3SegmentDescription,
		segments.K4SegmentDescription, segments.L1SegmentDescription, segments.N1SegmentDescription:
		if appendage, ok := f.Appendages[description]; ok {
			return appendage
		}
	}
	return nil
}

// GetSegment returns multiple segments like as j1, j2
func (f *fileInstance) GetListSegments(description string) []segments.Segment {
	switch description {
	case segments.J1SegmentDescription:
		return f.J1Segments
	case segments.J2SegmentDescription:
		return f.J2Segments
	}
	return nil
}

// GeneratorTrailer returns trailer segment that created automatically
func (f *fileInstance) GeneratorTrailer() (segments.Segment, error) {
	return nil, nil
}

// UnmarshalJSON parses a JSON blob
func (f *fileInstance) UnmarshalJSON(data []byte) error {
	f.reset()

	var dummy map[string]interface{}
	err := json.Unmarshal(data, &dummy)
	if err != nil {
		return nil
	}

	for key, record := range dummy {
		buf, err := json.Marshal(record)
		if err != nil {
			return err
		}

		switch key {
		case segments.BaseSegmentDescription:
			err = json.Unmarshal(buf, f.Base)
		case segments.HeaderRecordDescription:
			err = json.Unmarshal(buf, f.Header)
		case segments.TrailerRecordDescription:
			err = json.Unmarshal(buf, f.Trailer)
		case segments.J1SegmentDescription, segments.J2SegmentDescription:
			var list []interface{}
			err := json.Unmarshal(buf, &list)
			if err != nil {
				return nil
			}
			for _, subSegment := range list {
				subBuf, err := json.Marshal(subSegment)
				if err != nil {
					return err
				}
				err = f.unmarshalApplicableSegments(key, subBuf)
				if err != nil {
					return err
				}
			}
		default:
			err = f.unmarshalApplicableSegments(key, buf)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON returns JSON blob
func (f *fileInstance) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")

	value, err := json.Marshal(f.Header)
	if err != nil {
		return nil, err
	}
	buffer.WriteString(fmt.Sprintf("\"%s\":%s,", f.Header.Description(), string(value)))

	value, err = json.Marshal(f.Base)
	if err != nil {
		return nil, err
	}
	buffer.WriteString(fmt.Sprintf("\"%s\":%s,", f.Base.Description(), string(value)))

	if len(f.J1Segments) > 0 {
		value, err = json.Marshal(f.J1Segments)
		if err != nil {
			return nil, err
		}
		buffer.WriteString(fmt.Sprintf("\"%s\":%s,", segments.J1SegmentDescription, string(value)))
	}

	if len(f.J2Segments) > 0 {
		value, err = json.Marshal(f.J2Segments)
		if err != nil {
			return nil, err
		}
		buffer.WriteString(fmt.Sprintf("\"%s\":%s,", segments.J2SegmentDescription, string(value)))
	}

	for key, value := range f.Appendages {
		jsonValue, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		buffer.WriteString(fmt.Sprintf("\"%s\":%s,", key, string(jsonValue)))
	}

	trailerValue, err := json.Marshal(f.Trailer)
	if err != nil {
		return nil, err
	}
	buffer.WriteString(fmt.Sprintf("\"%s\":%s", f.Trailer.Description(), string(trailerValue)))

	buffer.WriteString("}")

	return buffer.Bytes(), nil
}

// Validate performs some checks on the file and returns an error if not Validated
func (f *fileInstance) Validate() error {
	return nil
}

// Parse attempts to initialize a *File object assuming the input is valid raw data.
func (f *fileInstance) Parse(record string) error {
	f.reset()
	offset := 0

	// Header Record
	hread, err := f.Header.Parse(record)
	if err != nil {
		return err
	}
	offset += hread

	// Data Record
	dread, err := f.Base.Parse(record[offset:])
	if err != nil {
		return err
	}

	positionSegments := f.Base.Length()
	if f.Base.BlockSize() > 0 {
		positionSegments += 4
	}
	read, err := f.readApplicableSegments(record[offset+positionSegments:])
	if err != nil {
		return err
	}

	if read+positionSegments > dread {
		return utils.NewErrFileParse()
	}
	offset += dread

	// Trailer Record
	tread, err := f.Trailer.Parse(record[offset:])
	if err != nil {
		return err
	}
	offset += tread

	if offset != len(record) {
		return utils.NewErrFileParse()
	}

	return nil
}

// String writes the File struct to raw string.
func (f *fileInstance) String() string {
	return ""
}

func (f *fileInstance) readApplicableSegments(record string) (int, error) {
	var segment segments.Segment
	offset := 0

	for offset < len(record) {
		switch record[offset : offset+2] {
		case segments.J1SegmentIdentifier:
			segment = segments.NewJ1Segment()
		case segments.J2SegmentIdentifier:
			segment = segments.NewJ2Segment()
		case segments.K1SegmentIdentifier:
			segment = segments.NewK1Segment()
		case segments.K2SegmentIdentifier:
			segment = segments.NewK2Segment()
		case segments.K3SegmentIdentifier:
			segment = segments.NewK3Segment()
		case segments.K4SegmentIdentifier:
			segment = segments.NewK4Segment()
		case segments.L1SegmentIdentifier:
			segment = segments.NewL1Segment()
		case segments.N1SegmentIdentifier:
			segment = segments.NewN1Segment()
		default:
			return offset, nil
		}
		read, err := segment.Parse(record[offset:])
		if err != nil {
			return 0, err
		}
		err = f.AddApplicableSegment(segment)
		if err != nil {
			return 0, err
		}
		offset += read
	}

	return offset, nil
}

func (f *fileInstance) unmarshalApplicableSegments(description string, data []byte) error {
	var segment segments.Segment

	switch description {
	case segments.J1SegmentDescription:
		segment = segments.NewJ1Segment()
	case segments.J2SegmentDescription:
		segment = segments.NewJ2Segment()
	case segments.K1SegmentDescription:
		segment = segments.NewK1Segment()
	case segments.K2SegmentDescription:
		segment = segments.NewK2Segment()
	case segments.K3SegmentDescription:
		segment = segments.NewK3Segment()
	case segments.K4SegmentDescription:
		segment = segments.NewK4Segment()
	case segments.L1SegmentDescription:
		segment = segments.NewL1Segment()
	case segments.N1SegmentDescription:
		segment = segments.NewN1Segment()
	default:
		return nil
	}

	err := json.Unmarshal(data, segment)
	if err != nil {
		return err
	}
	return f.AddApplicableSegment(segment)
}

func (f *fileInstance) reset() {
	f.Appendages = make(map[string]segments.Segment)
	f.J1Segments = []segments.Segment{}
	f.J2Segments = []segments.Segment{}
}
