// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

type Segment interface {
	Parse(record string) error
	String() string
	Validate() error
}
