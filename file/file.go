package file

import (
	"github.com/moov-io/metro2/segments"
	"github.com/moov-io/metro2/utils"
)

// General file interface
type File interface {
	SetBlock(segments.Segment) error
	AddApplicableSegment(segments.Segment) error
	GetSegment(string) segments.Segment
	GetListSegments(string) []segments.Segment
	GeneratorTrailer() (segments.Segment, error)
	UnmarshalJSON(p []byte) error
	MarshalJSON() ([]byte, error)

	Parse(record string) error
	String() string
	Validate() error
}

// NewFile constructs a file template.
func NewFile(format string) (File, error) {
	switch format {
	case CharacterFileFormat, PackedFileFormat:
		return &fileInstance{
			format:     CharacterFileFormat,
			Header:     segments.NewHeaderRecord(),
			Base:       segments.NewBaseSegment(),
			Appendages: make(map[string]segments.Segment),
			Trailer:    segments.NewTrailerRecord(),
		}, nil
	}
	return nil, utils.NewErrValidFileFormat(format)
}
