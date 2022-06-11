// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package utils

const (
	// Character Format (Fixed or variable blocked)
	CharacterFileFormat = "character"
	// Packed Format (Variable blocked)
	PackedFileFormat = "packed"
	// Name of header record
	HeaderRecordName = "header"
	// Name of trailer record
	TrailerRecordName = "trailer"
	// Name of data record
	DataRecordName = "data"
	// Length of packed record
	packedRecordLength = 366
	// Header identifier
	headerIdentifier = "HEADER"
	// Trailer identifier
	Trailerdentifier = "TRAILER"
	// Json format
	MessageJsonFormat = "json"
	// Metro format
	MessageMetroFormat = "metro"

	// Logging formats
	ColorRed   = "\033[31m"
	ColorGreen = "\033[32m"
	ColorBlue  = "\033[34m"
	ColorCyan  = "\033[36m"
)
