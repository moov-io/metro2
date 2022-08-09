// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package fuzzreader

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestCorpusSymlinks(t *testing.T) {
	// avoid symbolic link error on windows
	if runtime.GOOS == "windows" {
		t.Skip()
	}
	fds, err := os.ReadDir("corpus")
	if err != nil {
		t.Fatal(err)
	}
	if len(fds) == 0 {
		t.Fatal("no file descriptors found in corpus/")
	}

	for i := range fds {
		if fds[i].Type()&os.ModeSymlink != 0 {
			if path, err := os.Readlink(filepath.Join("corpus", fds[i].Name())); err != nil {
				t.Errorf("broken symlink: %v", err)
			} else {
				if _, err := os.Stat(filepath.Join("corpus", path)); err != nil {
					t.Errorf("broken symlink: %v", err)
				}
			}
		} else {
			t.Errorf("%s isn't a symlink, move outside corpus/ and symlink into directory", fds[i].Name())
		}
	}
}

func TestFuzzWithValidData(t *testing.T) {

	validFileSamples := []string{
		"packed_file.json",
		"packed_file.dat",
		"unpacked_fixed_file.dat",
		"unpacked_fixed_file.json",
		"unpacked_variable_file.dat",
		"unpacked_variable_file.json",
	}
	for _, sample := range validFileSamples {
		byteData, err := os.ReadFile(filepath.Join("..", "..", "test", "testdata", sample))
		if err != nil {
			t.Fatal(err)
		}

		if ret := Fuzz(byteData); ret != 1 {
			t.Errorf("Expected value is 1 (got %v)", ret)
		}
	}

	segmentSamples := []string{
		"base_segment.dat",
		"base_segment.json",
		"header_record.dat",
		"packed_base_segment.dat",
	}
	for _, sample := range segmentSamples {
		byteData, err := os.ReadFile(filepath.Join("..", "..", "test", "testdata", sample))
		if err != nil {
			t.Fatal(err)
		}

		if ret := Fuzz(byteData); ret != 0 {
			t.Errorf("Expected value is 0 (got %v)", ret)
		}
	}

}
