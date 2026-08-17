// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package utils

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// TestTime_MarshalJSON_Zero verifies that a zero-value Time (representing a
// not-available/zero-filled date field per the CDIA Metro2 spec) marshals to
// an empty JSON string rather than Go's zero time.Time value ("0001-01-01T00:00:00Z").
//
// See: https://github.com/moov-io/metro2/issues/196
func TestTime_MarshalJSON_Zero(t *testing.T) {
	var zero Time

	out, err := json.Marshal(zero)
	require.NoError(t, err)
	require.Equal(t, `""`, string(out))
}

// TestTime_MarshalJSON_NonZero verifies that a populated Time still marshals
// using the existing RFC3339 format, unaffected by the zero-value fix.
func TestTime_MarshalJSON_NonZero(t *testing.T) {
	tm := Time(time.Date(2020, time.January, 2, 15, 4, 5, 0, time.UTC))

	out, err := json.Marshal(tm)
	require.NoError(t, err)
	require.Equal(t, `"2020-01-02T15:04:05Z"`, string(out))
}

// TestTime_UnmarshalJSON_EmptyString verifies that an empty JSON string round
// trips back into a zero-value Time without erroring, so that Marshal/Unmarshal
// remain symmetric for not-available date fields.
func TestTime_UnmarshalJSON_EmptyString(t *testing.T) {
	var tm Time
	err := json.Unmarshal([]byte(`""`), &tm)
	require.NoError(t, err)
	require.True(t, tm.IsZero())
}
