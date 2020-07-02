// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"testing"
)

func TestValidator__upperAlpha(t *testing.T) {
	v := &validator{}
	if err := v.isUpperAlphanumeric("ab91"); err == nil {
		t.Error("expected error")
	}
}
