// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package fuzz

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/metro2/pkg/file"
	"github.com/moov-io/metro2/pkg/utils"
)

func FuzzReader(f *testing.F) {
	populateCorpus(f)

	f.Fuzz(func(t *testing.T, contents string) {
		f, _ := file.NewFileFromReader(strings.NewReader(contents))
		if f == nil {
			return
		}

		if record, _ := f.GetRecord(utils.HeaderRecordName); record == nil {
			return
		}
		if record, _ := f.GetRecord(utils.TrailerRecordName); record == nil {
			return
		}

		if records := f.GetDataRecords(); len(records) == 0 {
			return
		}

		f.Validate()
	})
}

func populateCorpus(f *testing.F) {
	f.Helper()

	err := filepath.Walk(filepath.Join("..", "testdata"), func(path string, info fs.FileInfo, _ error) error {
		if info.IsDir() {
			return nil
		}

		bs, err := os.ReadFile(path)
		if err != nil {
			f.Fatal(err)
		}
		f.Add(string(bs))
		return nil
	})
	if err != nil {
		f.Fatal(err)
	}
}
