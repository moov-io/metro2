package utils

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Time time.Time

// UnmarshalJSON Parses the json string in the custom format
func (ct *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := GetValidDate(s)
	*ct = Time(nt)
	return
}

// MarshalJSON writes a quoted string in the custom format
func (ct Time) MarshalJSON() ([]byte, error) {
	return []byte(ct.String()), nil
}

// String returns the time in the custom format
func (ct *Time) String() string {
	t := time.Time(*ct)

	return fmt.Sprintf("%q", t.Format(time.RFC3339))
}

func GetValidDate(timeString string) (time.Time, error) {

	date, err := ParseDate("01/02/2006", timeString)
	if err == nil {
		return date, nil
	}

	date, err = ParseDate("01/02/06", timeString)
	if err == nil {
		return date, nil
	}

	date, err = ParseDate("2006-01-02", timeString)
	if err == nil {
		return date, nil
	}

	date, err = ParseDate("1-02-2006", timeString)
	if err == nil {
		return date, nil
	}

	date, err = ParseDate("01-2-2006", timeString)
	if err == nil {
		return date, nil
	}

	date, err = ParseDate("1-2-2006", timeString)
	if err == nil {
		return date, nil
	}

	date, err = ParseDate("1/02/06", timeString)
	if err == nil {
		return date, nil
	}

	date, err = ParseDate("01/2/06", timeString)
	if err == nil {
		return date, nil
	}

	date, err = ParseDate("1/2/06", timeString)
	if err == nil {
		return date, nil
	}

	date, err = ParseDate("1/02/2006", timeString)
	if err == nil {
		return date, nil
	}

	date, err = ParseDate("01/2/2006", timeString)
	if err == nil {
		return date, nil
	}

	date, err = ParseDate("1/2/2006", timeString)
	if err == nil {
		return date, nil
	}

	date, err = ParseDate("02/01/2006", timeString)
	if err == nil {
		return date, nil
	}

	// "2006-01-02T15:04:05Z"
	date, err = ParseDate(time.RFC3339, timeString)
	if err == nil {
		return date, nil
	}

	//2020-06-01T14:49-06:00
	date, err = ParseDate("2006-01-02T15:04-07:00", timeString)
	if err == nil {
		return date, nil
	}

	date, err = ParseDate("2006-01", timeString)
	if err == nil {
		return date, nil
	}

	date, err = ParseDate("2006/01", timeString)
	if err == nil {
		return date, nil
	}

	date, err = ParseDate("Monday, January 02, 2006", timeString)
	if err == nil {
		return date, nil
	}

	date, err = ParseDate("1/2/2006 15:04", timeString)
	if err == nil {
		return date, nil
	}

	return time.Time{}, errors.New("Date is not parseable or not valid")
}

func ParseDate(dateLayout, timeString string) (time.Time, error) {
	if dateLayout == "" || timeString == "" {
		return time.Time{}, errors.New("Date is not parseable or not valid")
	}
	dateMarker := time.Now().AddDate(-200, 0, 0)

	// First try just using the provided date layout
	try1, err := time.Parse(dateLayout, timeString)
	if err == nil && try1.After(dateMarker) {
		return try1, nil
	}

	// Next try appending a common time layout
	timeLayout1 := "15:04:05 MST"
	try2, err := time.Parse(dateLayout+" "+timeLayout1, timeString)
	if err == nil && try2.After(dateMarker) {
		return try2, nil
	}

	// Next try appending another common time layout
	timeLayout2 := "03:04:05 PM"
	try3, err := time.Parse(dateLayout+" "+timeLayout2, timeString)
	if err == nil && try3.After(dateMarker) {
		return try3, nil
	}

	// If the time string has a space, split the time string so we can
	// check if the timeString has a null time at the end throwing off the parse
	if bytes.Contains([]byte(timeString), []byte(" ")) {
		pieces := strings.Split(timeString, " ")
		if len(pieces) > 0 {
			try4, err := time.Parse(dateLayout, strings.TrimSpace(pieces[0]))
			if err == nil && try4.After(dateMarker) {
				return try4, nil
			}
		}
	}

	return time.Time{}, errors.New("Date is not parseable or not valid")
}
