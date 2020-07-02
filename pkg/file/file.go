// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"encoding/json"

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
	case CharacterFileFormat:
		return &fileInstance{
			format:  CharacterFileFormat,
			Header:  lib.NewHeaderRecord(),
			Trailer: lib.NewTrailerRecord(),
		}, nil
	case PackedFileFormat:
		return &fileInstance{
			format:  PackedFileFormat,
			Header:  lib.NewPackedHeaderRecord(),
			Trailer: lib.NewPackedTrailerRecord(),
		}, nil
	}
	return nil, utils.NewErrValidFileFormat(format)
}

// CreateFile
func CreateFile(buf []byte) (File, error) {
	fileFormat, dataType, err := getFileInformation(buf)
	if err != nil {
		return nil, err
	}
	f, err := NewFile(*fileFormat)
	if err != nil {
		return nil, err
	}
	if *dataType == JsonData {
		err = json.Unmarshal(buf, f)
	} else {
		err = f.Parse(string(buf))
	}
	return f, err
}

func getFileInformation(buf []byte) (*string, *string, error) {
	fileFormat := CharacterFileFormat
	dataType := JsonData
	dummy := &dummyFile{}
	err := json.Unmarshal(buf, dummy)
	if err != nil {
		if !utils.IsMetroFile(string(buf)) {
			return nil, nil, utils.ErrValidField
		}
		dataType = MetroData
		if utils.IsVariableLength(string(buf)) {
			fileFormat = PackedFileFormat
		}
	} else {
		if dummy.Header == nil {
			return nil, nil, utils.ErrValidField
		}
		if dummy.Header.BlockDescriptorWord > 0 {
			fileFormat = PackedFileFormat
		}
	}
	return &fileFormat, &dataType, nil
}
