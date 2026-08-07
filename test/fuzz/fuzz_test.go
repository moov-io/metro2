// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package fuzz

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/metro2/pkg/file"
	"github.com/moov-io/metro2/pkg/utils"
)

func FuzzReader(f *testing.F) {
	populateCorpus(f, false)

	f.Fuzz(func(t *testing.T, contents string) {
		if len(contents) > 1<<20 {
			t.Skip()
		}

		parsed, err := file.NewFileFromReader(strings.NewReader(contents))
		if err != nil || parsed == nil {
			return
		}

		if record, _ := parsed.GetRecord(utils.HeaderRecordName); record == nil {
			return
		}
		if record, _ := parsed.GetRecord(utils.TrailerRecordName); record == nil {
			return
		}

		_ = parsed.Validate()
		_ = parsed.String(true)
		_ = parsed.String(false)
		_ = parsed.Bytes()
	})
}

func FuzzJSON(f *testing.F) {
	populateCorpus(f, true)

	f.Fuzz(func(t *testing.T, contents string) {
		if len(contents) > 1<<20 {
			t.Skip()
		}

		// Try both packed and character formats
		for _, format := range []string{utils.CharacterFileFormat, utils.PackedFileFormat} {
			fl, err := file.NewFile(format)
			if err != nil || fl == nil {
				continue
			}
			if err := json.Unmarshal([]byte(contents), fl); err != nil {
				continue
			}
			_ = fl.Validate()
			_ = fl.String(true)
		}
	})
}

func populateCorpus(f *testing.F, jsonOnly bool) {
	f.Helper()

	f.Add("")
	f.Add("{}")

	err := filepath.Walk(filepath.Join("..", "testdata"), func(path string, info fs.FileInfo, _ error) error {
		if info == nil || info.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		if jsonOnly && ext != ".json" {
			return nil
		}
		if !jsonOnly && ext == ".json" {
			// still useful for NewFileFromReader which may detect format
		}

		bs, err := os.ReadFile(path)
		if err != nil {
			f.Fatal(err)
		}
		if len(bs) > 256*1024 {
			return nil
		}
		f.Add(string(bs))
		return nil
	})
	if err != nil {
		f.Fatal(err)
	}
}
