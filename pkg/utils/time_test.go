// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package utils

import (
	"bytes"
	"testing"
)

func TestMarshalJSONIsZeroDate(t *testing.T) {
	ct := new(Time) // create zero time
	str, _ := ct.MarshalJSON()
	if !bytes.Equal(str, []byte("\"\"")) {
		t.Error("Expected empty string but received: \"" + string(str) + "\"\n")
	}
}
