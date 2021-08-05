// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"encoding/json"

	"github.com/moov-io/base/log"
	"github.com/moov-io/metro2/pkg/lib"
	"github.com/moov-io/metro2/pkg/utils"
)

// General file interface
type File interface {
	SetRecord(lib.Record) error
	AddDataRecord(lib.Record) error
	GetRecord(string) (lib.Record, error)
	GetDataRecords() []lib.Record
	GeneratorTrailer() (lib.Record, error)

	Parse(record string) error
	String() string
	Validate() error
}

type headerInformation struct {
	BlockDescriptorWord  int    `json:"blockDescriptorWord"`
	RecordDescriptorWord int    `json:"recordDescriptorWord"`
	RecordIdentifier     string `json:"recordIdentifier"`
}

type dummyFile struct {
	Header *headerInformation `json:"header"`
}

// NewFile constructs a file template.
func NewFile(format string) (File, error) {
	switch format {
	case utils.CharacterFileFormat:
		return &fileInstance{
			logger:  log.NewDefaultLogger(),
			format:  utils.CharacterFileFormat,
			Header:  lib.NewHeaderRecord(),
			Trailer: lib.NewTrailerRecord(),
		}, nil
	case utils.PackedFileFormat:
		return &fileInstance{
			logger:  log.NewDefaultLogger(),
			format:  utils.PackedFileFormat,
			Header:  lib.NewPackedHeaderRecord(),
			Trailer: lib.NewPackedTrailerRecord(),
		}, nil
	}
	return nil, utils.NewErrInvalidSegment(format)
}

// CreateFile attempts to parse raw metro2 file contents
func CreateFile(buf []byte) (File, error) {
	fileFormat, dataType, err := getFileInformation(buf)
	if err != nil {
		return nil, err
	}
	f, _ := NewFile(*fileFormat)
	if *dataType == utils.MessageJsonFormat {
		err = json.Unmarshal(buf, f)
	} else {
		err = f.Parse(string(buf))
	}
	return f, err
}

func getFileInformation(buf []byte) (*string, *string, error) {
	fileFormat := utils.CharacterFileFormat
	dataType := utils.MessageJsonFormat
	dummy := &dummyFile{}
	err := json.Unmarshal(buf, dummy)
	if err != nil {
		if !utils.IsMetroFile(string(buf)) {
			return nil, nil, utils.ErrInvalidMetroFile
		}
		dataType = utils.MessageMetroFormat
		if utils.IsPacked(string(buf)) {
			fileFormat = utils.PackedFileFormat
		}
	} else {
		if dummy.Header == nil {
			return nil, nil, utils.ErrNonHeaderRecord
		}
		if dummy.Header.RecordDescriptorWord == lib.UnpackedRecordLength {
			fileFormat = utils.CharacterFileFormat
		} else if dummy.Header.BlockDescriptorWord > 0 {
			fileFormat = utils.PackedFileFormat
		}
	}
	return &fileFormat, &dataType, nil
}
