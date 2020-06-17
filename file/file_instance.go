package file

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/moov-io/metro2/segments"
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

func (f *fileInstance) SetSegment(s segments.Segment) error {
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

func (f *fileInstance) GetSegments(description string) []segments.Segment {
	switch description {
	case segments.J1SegmentDescription:
		return f.J1Segments
	case segments.J2SegmentDescription:
		return f.J2Segments
	}
	return nil
}

func (f *fileInstance) GeneratorTrailer() (segments.Segment, error) {
	return nil, nil
}

func (f *fileInstance) UnmarshalJSON(p []byte) error {
	return nil
}

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

func (f *fileInstance) Validate() error {
	return nil
}

func (f *fileInstance) Parse(record string) error {
	return nil
}

func (f *fileInstance) String() string {
	return ""
}
