// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"io"
	"strconv"

	"github.com/moov-io/base/log"
	"github.com/moov-io/metro2/pkg/lib"
	"github.com/moov-io/metro2/pkg/utils"
)

// General file interface
type File interface {
	GetType() string
	SetType(string) error
	SetRecord(lib.Record) error
	AddDataRecord(lib.Record) error
	GetRecord(string) (lib.Record, error)
	GetDataRecords() []lib.Record
	GeneratorTrailer() (lib.Record, error)

	Parse(record []byte) error
	Bytes() []byte
	String(newline bool) string
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

// CreateFile attempts to parse raw metro2 file or json file
func CreateFile(buf []byte) (File, error) {
	fileFormat, dataType, err := getFileInformation(buf)
	if err != nil {
		return nil, err
	}
	f, _ := NewFile(*fileFormat)
	if *dataType == utils.MessageJsonFormat {
		err = json.Unmarshal(buf, f)
	} else {
		err = f.Parse(buf)
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

// Reader reads records from a metro2 encoded file.
type Reader struct {
	// r handles the IO.Reader sent to be parser.
	scanner *bufio.Scanner
	// file is ach.file model being built as r is parsed.
	File File
	// line is the current line being parsed from the input r
	line []byte
	// recordName holds the current record name being parsed.
	recordName string
}

// Read reads each record of the metro file
func (r *Reader) Read() (File, error) {

	f := r.File.(*fileInstance)

	f.Bases = []lib.Record{}

	// read through the entire file
	if r.scanner.Scan() {
		r.line = r.scanner.Bytes()

		// getting file type
		fileFormat, _, err := getFileInformation([]byte(r.line))
		if err != nil {
			return nil, err
		}

		f.SetType(*fileFormat)

		// Header Record
		_, err = f.Header.Parse(r.line)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, utils.NewErrInvalidSegment("header")
	}

	failedParse := false
	for r.scanner.Scan() {
		r.line = r.scanner.Bytes()

		var base lib.Record
		if f.format == utils.PackedFileFormat {
			base = lib.NewPackedBaseSegment()
		} else {
			base = lib.NewBaseSegment()
		}

		_, err := base.Parse(r.line)
		if err != nil {
			failedParse = true
			break
		}
		f.Bases = append(f.Bases, base)
	}

	if !failedParse {
		// read new line
		if r.scanner.Scan() {
			r.line = r.scanner.Bytes()
		} else {
			return nil, utils.NewErrInvalidSegment("trailer")
		}
	}

	_, err := f.Trailer.Parse(r.line)
	if err != nil {
		return nil, err
	}

	return r.File, nil
}

// NewReader returns a new metro reader that reads from io reader.
func NewReader(r io.Reader) *Reader {
	f, _ := NewFile(utils.CharacterFileFormat)
	reader := &Reader{
		File:    f,
		scanner: bufio.NewScanner(r),
	}

	reader.scanner.Split(scanRecord)

	return reader
}

//scanRecord allows reader to split metro file by each record
func scanRecord(data []byte, atEOF bool) (advance int, token []byte, err error) {

	getStripedLength := func() int {
		return len(bytes.ReplaceAll(bytes.ReplaceAll(data, []byte("\r\n"), nil), []byte("\n"), nil))
	}

	getStripedData := func(size int) (int, []byte, error) {
		for i := size; i <= len(data); i++ {
			converted := bytes.ReplaceAll(bytes.ReplaceAll(data[:i], []byte("\r\n"), nil), []byte("\n"), nil)
			if len(converted) == size {
				return i, converted, nil
			}
		}
		return 0, nil, io.ErrNoProgress
	}

	length := getStripedLength()

	if atEOF && length == 0 {
		return 0, nil, nil
	} else if length < 4 && atEOF {
		// we ran out of bytes and we're at the end of the file
		return 0, nil, io.ErrUnexpectedEOF
	} else if length < 4 {
		// we need at least the control bytes
		return 0, nil, nil
	}

	_, bdw, _ := getStripedData(4)
	// trying to read for unpacked format
	size, readErr := strconv.ParseInt(string(bdw), 10, 64)
	if readErr == nil {
		if size < lib.UnpackedRecordLength {
			return 0, nil, io.ErrNoProgress
		}
	} else {
		// trying to read for packed format
		size = int64(binary.BigEndian.Uint16(bdw))
		if size < lib.PackedRecordLength {
			return 0, nil, io.ErrNoProgress
		}
	}

	if int(size) <= length {
		// return line while accounting for control bytes
		return getStripedData(int(size))
	} else if int(size) > length && atEOF {
		// we need more data, but there is no more data to read
		return 0, nil, io.ErrUnexpectedEOF
	}

	// request more data.
	return 0, nil, nil
}
