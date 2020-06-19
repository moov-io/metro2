package file

import (
	"encoding/json"
	"sort"
	"strings"

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

	dummy := make(map[string]interface{})
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
		case segments.BaseSegmentDescription, segments.PackedBaseSegmentDescription:
			err = json.Unmarshal(buf, f.Base)
		case segments.HeaderRecordDescription, segments.PackedHeaderRecordDescription:
			err = json.Unmarshal(buf, f.Header)
		case segments.TrailerRecordDescription, segments.PackedTrailerRecordDescription:
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

type fileInstanceJson struct {
	Header        segments.Segment   `json:"headerCharacter,omitempty"`
	Base          segments.Segment   `json:"baseCharacter,omitempty"`
	PackedHeader  segments.Segment   `json:"headerPacked,omitempty"`
	PackedBase    segments.Segment   `json:"basePacked,omitempty"`
	J1Segments    []segments.Segment `json:"j1,omitempty"`
	J2Segments    []segments.Segment `json:"j2,omitempty"`
	K1Segments    segments.Segment   `json:"k1,omitempty"`
	K2Segments    segments.Segment   `json:"k2,omitempty"`
	K3Segments    segments.Segment   `json:"k3,omitempty"`
	K4Segments    segments.Segment   `json:"k4,omitempty"`
	L1Segments    segments.Segment   `json:"l1,omitempty"`
	N1Segments    segments.Segment   `json:"n1,omitempty"`
	Trailer       segments.Segment   `json:"trailer,omitempty"`
	PackedTrailer segments.Segment   `json:"trailerPacked,omitempty"`
}

// MarshalJSON returns JSON blob
func (f *fileInstance) MarshalJSON() ([]byte, error) {
	dummy := fileInstanceJson{}

	if f.Header.Description() == segments.HeaderRecordDescription {
		dummy.Header = f.Header
	} else {
		dummy.PackedHeader = f.Header
	}
	if f.Base.Description() == segments.BaseSegmentDescription {
		dummy.Base = f.Base
	} else {
		dummy.PackedBase = f.Base
	}
	if f.Trailer.Description() == segments.TrailerRecordDescription {
		dummy.Trailer = f.Trailer
	} else {
		dummy.PackedTrailer = f.Trailer
	}
	dummy.J1Segments = f.J1Segments
	dummy.J2Segments = f.J2Segments

	for key, sub := range f.Appendages {
		switch key {
		case segments.K1SegmentDescription:
			dummy.K1Segments = sub
		case segments.K2SegmentDescription:
			dummy.K2Segments = sub
		case segments.K3SegmentDescription:
			dummy.K3Segments = sub
		case segments.K4SegmentDescription:
			dummy.K4Segments = sub
		case segments.L1SegmentDescription:
			dummy.L1Segments = sub
		case segments.N1SegmentDescription:
			dummy.N1Segments = sub
		}
	}
	return json.Marshal(dummy)
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
	var buf strings.Builder

	// Header Block
	header := f.Header.String()

	// Data Block
	base := f.Base.String()
	dataBlockSize := len(base)
	baseSize := f.Base.Length()
	if f.Base.BlockSize() > 0 {
		baseSize += 4
	}
	base = base[:baseSize]

	for _, sub := range f.J1Segments {
		base += sub.String()
	}
	for _, sub := range f.J2Segments {
		base += sub.String()
	}

	if len(f.Appendages) > 0 {
		var keys []string
		for key := range f.Appendages {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			base += f.Appendages[key].String()
		}
	}

	if dataBlockSize > len(base) {
		base += strings.Repeat(" ", dataBlockSize-len(base))
	}

	// Trailer Block
	trailer := f.Trailer.String()

	buf.Grow(len(header) + len(base) + len(trailer))
	buf.WriteString(header)
	buf.WriteString(base)
	buf.WriteString(trailer)

	return buf.String()
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
