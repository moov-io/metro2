// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"github.com/moov-io/metro2/lib"
	"github.com/moov-io/metro2/utils"
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
