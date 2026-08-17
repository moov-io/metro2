// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import "testing"

func TestClampedBlockSize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		requested int
		want      int
	}{
		{name: "zero", requested: 0, want: 0},
		{name: "negative", requested: -1, want: 0},
		{name: "standard unpacked", requested: UnpackedRecordLength, want: UnpackedRecordLength},
		{name: "standard packed", requested: PackedRecordLength, want: PackedRecordLength},
		{name: "max allowed", requested: maxRecordPad, want: maxRecordPad},
		{name: "just over max", requested: maxRecordPad + 1, want: 0},
		{name: "fuzz rdw", requested: 5175555555, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := clampedBlockSize(tt.requested); got != tt.want {
				t.Fatalf("clampedBlockSize(%d) = %d, want %d", tt.requested, got, tt.want)
			}
		})
	}
}

func TestHeaderRecordStringHugeDescriptorWord(t *testing.T) {
	t.Parallel()

	// Reproduces FuzzReader crash 1b557a3a811c2d1b: JSON with a multi-GB
	// recordDescriptorWord used to make String() allocate until the process died.
	record := &HeaderRecord{
		RecordDescriptorWord: 5175555555,
		RecordIdentifier:     HeaderIdentifier,
	}
	out := record.String()
	if len(out) > maxRecordPad {
		t.Fatalf("String() allocated %d bytes from huge descriptor word", len(out))
	}
	if len(out) == 0 {
		t.Fatal("String() returned empty output")
	}
}

func TestTrailerRecordStringHugeDescriptorWord(t *testing.T) {
	t.Parallel()

	record := &TrailerRecord{
		RecordDescriptorWord: 5175555555,
		RecordIdentifier:     TrailerIdentifier,
	}
	out := record.String()
	if len(out) > maxRecordPad {
		t.Fatalf("String() allocated %d bytes from huge descriptor word", len(out))
	}
}

func TestBaseSegmentStringHugeDescriptorWord(t *testing.T) {
	t.Parallel()

	record := &BaseSegment{
		RecordDescriptorWord: 5175555555,
	}
	out := record.String()
	if len(out) > maxRecordPad {
		t.Fatalf("String() allocated %d bytes from huge descriptor word", len(out))
	}
}

func TestHeaderRecordStringNegativeDescriptorWord(t *testing.T) {
	t.Parallel()

	record := &HeaderRecord{
		RecordDescriptorWord: -1,
		RecordIdentifier:     HeaderIdentifier,
	}
	// Must not panic on strings.Builder.Grow / strings.Repeat
	_ = record.String()
}
