// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"bytes"
	"encoding/json"
	"errors"
	"reflect"
	"strings"
	"unicode"

	"github.com/moov-io/base/log"
	"github.com/moov-io/metro2/pkg/lib"
	"github.com/moov-io/metro2/pkg/utils"
)

var _ File = (*fileInstance)(nil)

// File contains the structures of a parsed metro 2 file.
type fileInstance struct {
	Header  lib.Record   `json:"header"`
	Bases   []lib.Record `json:"data"`
	Trailer lib.Record   `json:"trailer"`

	format string
	logger log.Logger
}

// SetRecord can set block record like as header, trailer
func (f *fileInstance) SetRecord(r lib.Record) error {
	err := r.Validate()
	if err != nil {
		return err
	}

	if (f.format == utils.PackedFileFormat && r.Name() == lib.PackedHeaderRecordName) ||
		(f.format == utils.CharacterFileFormat && r.Name() == lib.HeaderRecordName) {
		f.Header = r
	} else if (f.format == utils.PackedFileFormat && r.Name() == lib.PackedTrailerRecordName) ||
		(f.format == utils.CharacterFileFormat && r.Name() == lib.TrailerRecordName) {
		f.Trailer = r
	} else {
		return utils.NewErrInvalidRecord(r.Name())
	}

	return nil
}

// AddDataRecord can append data record
func (f *fileInstance) AddDataRecord(r lib.Record) error {
	err := r.Validate()
	if err != nil {
		return err
	}

	if f.format == utils.PackedFileFormat && r.Name() == lib.PackedBaseSegmentName {
		f.Bases = append(f.Bases, r)
	} else if f.format == utils.CharacterFileFormat && r.Name() == lib.BaseSegmentName {
		f.Bases = append(f.Bases, r)
	} else {
		return utils.NewErrInvalidRecord(r.Name())
	}

	return nil
}

// GetRecord returns single record like as header, trailer.
func (f *fileInstance) GetRecord(name string) (lib.Record, error) {
	switch name {
	case utils.HeaderRecordName:
		return f.Header, nil
	case utils.TrailerRecordName:
		return f.Trailer, nil
	default:
		return nil, utils.NewErrInvalidRecord(name)
	}
}

// GetDataRecords returns data records
func (f *fileInstance) GetDataRecords() []lib.Record {
	return f.Bases
}

// GeneratorTrailer returns trailer segment that created automatically
func (f *fileInstance) GeneratorTrailer() (lib.Record, error) {
	var trailer lib.Record
	var information *lib.TrailerInformation
	var err error

	if f.format == utils.PackedFileFormat {
		trailer = lib.NewPackedTrailerRecord()
		information, err = f.generatorPackedTrailer()
		if err != nil {
			return nil, err
		}
	} else {
		trailer = lib.NewTrailerRecord()
		information, err = f.generatorTrailer()
		if err != nil {
			return nil, err
		}
	}

	fromFields := reflect.ValueOf(information).Elem()
	toFields := reflect.ValueOf(trailer).Elem()
	for i := 0; i < fromFields.NumField(); i++ {
		fieldName := fromFields.Type().Field(i).Name
		fromField := fromFields.FieldByName(fieldName)
		toField := toFields.FieldByName(fieldName)
		if fromField.IsValid() && toField.CanSet() {
			toField.Set(fromField)
		}
	}

	if f.format == utils.PackedFileFormat {
		if segment, ok := trailer.(*lib.PackedTrailerRecord); ok {
			segment.RecordDescriptorWord = lib.PackedRecordLength
			segment.RecordIdentifier = lib.TrailerIdentifier
		}
	} else {
		if segment, ok := trailer.(*lib.TrailerRecord); ok {
			segment.RecordDescriptorWord = lib.UnpackedRecordLength
			segment.RecordIdentifier = lib.TrailerIdentifier
		}
	}

	return trailer, nil
}

// Validate performs some checks on the file and returns an error if not Validated
func (f *fileInstance) Validate() error {
	var err error

	if f.Header != nil {
		err = f.Header.Validate()
		if err != nil {
			return err
		}
	}

	if f.Trailer != nil {
		err = f.Trailer.Validate()
		if err != nil {
			return err
		}
	}

	for _, baseSegment := range f.Bases {
		err = baseSegment.Validate()
		if err != nil {
			return err
		}
	}

	var information *lib.TrailerInformation
	if f.format == utils.PackedFileFormat {
		information, err = f.generatorPackedTrailer()
		if err != nil {
			return err
		}
	} else {
		information, err = f.generatorTrailer()
		if err != nil {
			return err
		}
	}

	fromFields := reflect.ValueOf(information).Elem()
	toFields := reflect.ValueOf(f.Trailer).Elem()
	for i := 0; i < fromFields.NumField(); i++ {
		fieldName := fromFields.Type().Field(i).Name

		// skip local variable
		if !unicode.IsUpper([]rune(fieldName)[0]) {
			continue
		}
		switch fieldName {
		case "BlockDescriptorWord", "RecordDescriptorWord", "RecordIdentifier":
			continue
		}

		fromField := fromFields.FieldByName(fieldName)
		toField := toFields.FieldByName(fieldName)
		if !fromField.IsValid() || !toField.IsValid() {
			return utils.NewErrInvalidValueOfField(fieldName, "trailer record")
		}
		if fromField.Interface() != toField.Convert(fromField.Type()).Interface() {
			return utils.NewErrInvalidValueOfField(fieldName, "trailer record")
		}
	}
	return nil
}

// Parse attempts to initialize a *File object assuming the input is valid raw data.
func (f *fileInstance) Parse(record []byte) error {

	// remove new lines
	record = bytes.ReplaceAll(bytes.ReplaceAll(record, []byte("\r\n"), nil), []byte("\n"), nil)

	f.Bases = []lib.Record{}
	offset := 0

	// Header Record
	head, err := f.Header.Parse(record)
	if err != nil {
		return err
	}
	offset += head

	// Data Record
	for err == nil {
		var base lib.Record
		if f.format == utils.PackedFileFormat {
			base = lib.NewPackedBaseSegment()
		} else {
			base = lib.NewBaseSegment()
		}

		if offset <= 0 || len(record) <= offset {
			return utils.NewErrSegmentLength("base record")
		}

		read, err := base.Parse(record[offset:])
		if err != nil {
			break
		}
		f.Bases = append(f.Bases, base)
		offset += read
	}

	// Trailer Record
	if offset <= 0 || len(record) <= offset {
		return utils.NewErrSegmentLength("trailer record")
	}
	tread, err := f.Trailer.Parse(record[offset:])
	if err != nil {
		return err
	}
	offset += tread

	if offset != len(record) {
		return utils.NewErrFailedParsing()
	}

	return nil
}

// String writes the File struct to raw string.
func (f *fileInstance) String(isNewLine bool) string {
	var buf strings.Builder

	newLine := ""
	if isNewLine {
		newLine = "\n"
	}

	// Header Block
	header := f.Header.String() + newLine

	// Data Block
	data := ""
	for _, base := range f.Bases {
		data += base.String() + newLine
	}

	// Trailer Block
	trailer := f.Trailer.String()

	buf.Grow(len(header) + len(data) + len(trailer))
	buf.WriteString(header)
	buf.WriteString(data)
	buf.WriteString(trailer)

	return buf.String()
}

// Bytes return raw byte array
func (r *fileInstance) Bytes() []byte {
	return []byte(r.String(false))
}

// UnmarshalJSON parses a JSON blob
func (f *fileInstance) UnmarshalJSON(data []byte) error {
	dummy := make(map[string]interface{})
	err := json.Unmarshal(data, &dummy)
	if err != nil {
		return err
	}

	for name, record := range dummy {
		buf, err := json.Marshal(record)
		if err != nil {
			return err
		}

		switch name {
		case utils.HeaderRecordName:
			err = json.Unmarshal(buf, f.Header)
			if err != nil {
				f.logger.Error().LogErrorf(err.Error())
				return errors.New("Unable to parse input json file")
			}
		case utils.TrailerRecordName:
			err = json.Unmarshal(buf, f.Trailer)
			if err != nil {
				f.logger.Error().LogErrorf(err.Error())
				return errors.New("Unable to parse input json file")
			}
		case utils.DataRecordName:
			var list []interface{}
			err = json.Unmarshal(buf, &list)
			if err != nil {
				f.logger.Error().LogErrorf(err.Error())
				return errors.New("Unable to parse input json file")
			}
			for _, subSegment := range list {
				subBuf, err := json.Marshal(subSegment)
				if err != nil {
					return err
				}
				if f.format == utils.CharacterFileFormat {
					base := lib.NewBaseSegment()
					err = json.Unmarshal(subBuf, base)
					if err != nil {
						f.logger.Error().LogErrorf(err.Error())
						return errors.New("Unable to parse input json file")
					}
					f.Bases = append(f.Bases, base)
				} else {
					base := lib.NewPackedBaseSegment()
					err = json.Unmarshal(subBuf, base)
					if err != nil {
						f.logger.Error().LogErrorf(err.Error())
						return errors.New("Unable to parse input json file")
					}
					f.Bases = append(f.Bases, base)
				}
			}
		}
	}

	return nil
}

// GetType returns type of the file
func (f *fileInstance) GetType() string {
	return f.format
}

// GetType returns type of the file
func (f *fileInstance) SetType(newType string) error {
	if newType != utils.CharacterFileFormat && newType != utils.PackedFileFormat {
		return errors.New("invalid type")
	}

	if f.format == newType {
		return nil
	}

	if newType == utils.CharacterFileFormat {
		// convert header
		if f.Header != nil {
			newHeader := lib.HeaderRecord(*f.Header.(*lib.PackedHeaderRecord))
			f.Header = &newHeader
		}

		// convert bases
		var bases []lib.Record
		for _, _base := range f.Bases {
			newBase := lib.BaseSegment(*_base.(*lib.PackedBaseSegment))
			bases = append(bases, &newBase)
		}
		f.Bases = bases

		// convert trailer
		if f.Trailer != nil {
			newTrailer := lib.TrailerRecord(*f.Trailer.(*lib.PackedTrailerRecord))
			f.Trailer = &newTrailer
		}
	} else if newType == utils.PackedFileFormat {
		// convert header
		if f.Header != nil {
			newHeader := lib.PackedHeaderRecord(*f.Header.(*lib.HeaderRecord))
			f.Header = &newHeader
		}

		// convert bases
		var bases []lib.Record
		for _, _base := range f.Bases {
			newBase := lib.PackedBaseSegment(*_base.(*lib.BaseSegment))
			bases = append(bases, &newBase)
		}
		f.Bases = bases

		// convert trailer
		if f.Trailer != nil {
			newTrailer := lib.PackedTrailerRecord(*f.Trailer.(*lib.TrailerRecord))
			f.Trailer = &newTrailer
		}
	}

	f.format = newType
	return nil
}

func (f *fileInstance) generatorTrailer() (*lib.TrailerInformation, error) {
	trailer := &lib.TrailerInformation{}

	trailer.TotalBaseRecords = len(f.Bases)
	trailer.BlockCount = len(f.Bases) + 2
	for _, base := range f.Bases {
		base, ok := base.(*lib.BaseSegment)
		if !ok && base.Validate() != nil {
			return nil, utils.NewErrInvalidSegment(base.Name())
		}

		if isValidSocialSecurityNumber(base.SocialSecurityNumber) {
			trailer.TotalSocialNumbersAllSegments++
			trailer.TotalSocialNumbersBaseSegments++
		}

		if !base.DateBirth.IsZero() {
			trailer.TotalDatesBirthAllSegments++
			trailer.TotalDatesBirthBaseSegments++
		}

		if base.ECOACode == lib.ECOACodeZ {
			trailer.TotalECOACodeZ++
		}
		if base.TelephoneNumber > 0 {
			trailer.TotalTelephoneNumbersAllSegments++
		}
		f.statisticAccountStatus(base.AccountStatus, trailer)
		f.statisticBase(base, trailer)
	}

	return trailer, nil
}

func (f *fileInstance) generatorPackedTrailer() (*lib.TrailerInformation, error) {
	trailer := &lib.TrailerInformation{}
	trailer.TotalBaseRecords = len(f.Bases)
	trailer.BlockCount = len(f.Bases) + 2
	for _, base := range f.Bases {
		base, ok := base.(*lib.PackedBaseSegment)
		if !ok && base.Validate() != nil {
			return nil, utils.NewErrInvalidSegment(base.Name())
		}

		if isValidSocialSecurityNumber(base.SocialSecurityNumber) {
			trailer.TotalSocialNumbersAllSegments++
			trailer.TotalSocialNumbersBaseSegments++
		}

		if !base.DateBirth.IsZero() {
			trailer.TotalDatesBirthAllSegments++
			trailer.TotalDatesBirthBaseSegments++
		}

		if base.ECOACode == lib.ECOACodeZ {
			trailer.TotalECOACodeZ++
		}

		if base.TelephoneNumber > 0 {
			trailer.TotalTelephoneNumbersAllSegments++
		}

		f.statisticAccountStatus(base.AccountStatus, trailer)
		f.statisticPackedBase(base, trailer)
	}

	return trailer, nil
}

func (f *fileInstance) statisticAccountStatus(status string, info *lib.TrailerInformation) {
	switch status {
	case lib.AccountStatusDF:
		info.TotalStatusCodeDF++
	case lib.AccountStatusDA:
		info.TotalStatusCodeDA++
	case lib.AccountStatus05:
		info.TotalStatusCode05++
	case lib.AccountStatus11:
		info.TotalStatusCode11++
	case lib.AccountStatus13:
		info.TotalStatusCode13++
	case lib.AccountStatus61:
		info.TotalStatusCode61++
	case lib.AccountStatus62:
		info.TotalStatusCode62++
	case lib.AccountStatus63:
		info.TotalStatusCode63++
	case lib.AccountStatus64:
		info.TotalStatusCode64++
	case lib.AccountStatus65:
		info.TotalStatusCode65++
	case lib.AccountStatus71:
		info.TotalStatusCode71++
	case lib.AccountStatus78:
		info.TotalStatusCode78++
	case lib.AccountStatus80:
		info.TotalStatusCode80++
	case lib.AccountStatus82:
		info.TotalStatusCode82++
	case lib.AccountStatus83:
		info.TotalStatusCode83++
	case lib.AccountStatus84:
		info.TotalStatusCode84++
	case lib.AccountStatus88:
		info.TotalStatusCode88++
	case lib.AccountStatus89:
		info.TotalStatusCode89++
	case lib.AccountStatus93:
		info.TotalStatusCode93++
	case lib.AccountStatus94:
		info.TotalStatusCode94++
	case lib.AccountStatus95:
		info.TotalStatusCode95++
	case lib.AccountStatus96:
		info.TotalStatusCode96++
	case lib.AccountStatus97:
		info.TotalStatusCode97++
	}
}

func (f *fileInstance) statisticPackedBase(base *lib.PackedBaseSegment, trailer *lib.TrailerInformation) {
	for _, j1 := range base.GetSegments(lib.J1SegmentName) {
		sub := j1.(*lib.J1Segment)
		if sub.ECOACode == lib.ECOACodeZ {
			trailer.TotalECOACodeZ++
		}
		if sub.Validate() == nil {
			trailer.TotalConsumerSegmentsJ1++

			if isValidSocialSecurityNumber(sub.SocialSecurityNumber) {
				trailer.TotalSocialNumbersAllSegments++
				trailer.TotalSocialNumbersJ1Segments++
			}

			if !sub.DateBirth.IsZero() {
				trailer.TotalDatesBirthAllSegments++
				trailer.TotalDatesBirthJ1Segments++
			}

			if sub.TelephoneNumber > 0 {
				trailer.TotalTelephoneNumbersAllSegments++
			}
		}
	}
	for _, j2 := range base.GetSegments(lib.J2SegmentName) {
		sub := j2.(*lib.J2Segment)
		if sub.ECOACode == lib.ECOACodeZ {
			trailer.TotalECOACodeZ++
		}
		if sub.Validate() == nil {
			trailer.TotalConsumerSegmentsJ2++

			if isValidSocialSecurityNumber(sub.SocialSecurityNumber) {
				trailer.TotalSocialNumbersAllSegments++
				trailer.TotalSocialNumbersJ2Segments++
			}

			if !sub.DateBirth.IsZero() {
				trailer.TotalDatesBirthAllSegments++
				trailer.TotalDatesBirthJ2Segments++
			}

			if sub.TelephoneNumber > 0 {
				trailer.TotalTelephoneNumbersAllSegments++
			}
		}
	}
	for _, k1 := range base.GetSegments(lib.K1SegmentName) {
		sub := k1.(*lib.K1Segment)
		if len(sub.OriginalCreditorName) > 0 {
			trailer.TotalOriginalCreditorSegments++
		}
	}
	for _, k2 := range base.GetSegments(lib.K2SegmentName) {
		sub := k2.(*lib.K2Segment)
		if sub.PurchasedIndicator == lib.PurchasedIndicatorToName ||
			sub.PurchasedIndicator == lib.PurchasedIndicatorFromName {
			trailer.TotalPurchasedToSegments++
		}
	}
	for _, k3 := range base.GetSegments(lib.K3SegmentName) {
		sub := k3.(*lib.K3Segment)
		if sub.AgencyIdentifier == lib.AgencyIdentifierNotApplicable {
			trailer.TotalMortgageInformationSegments++
		}
	}
	for _, k4 := range base.GetSegments(lib.K4SegmentName) {
		sub := k4.(*lib.K4Segment)
		if sub.SpecializedPaymentIndicator == lib.SpecializedBalloonPayment ||
			sub.SpecializedPaymentIndicator == lib.SpecializedDeferredPayment {
			trailer.TotalPaymentInformationSegments++
		}
	}
	for _, l1 := range base.GetSegments(lib.L1SegmentName) {
		sub := l1.(*lib.L1Segment)
		if sub.ChangeIndicator == lib.ChangeIndicatorAccountNumber ||
			sub.ChangeIndicator == lib.ChangeIndicatorIdentificationNumber ||
			sub.ChangeIndicator == lib.ChangeIndicatorBothNumber {
			trailer.TotalChangeSegments++
		}
	}
	for _, n1 := range base.GetSegments(lib.N1SegmentName) {
		sub := n1.(*lib.N1Segment)
		if len(sub.EmployerName) > 0 {
			trailer.TotalEmploymentSegments++
		}
	}
}

func (f *fileInstance) statisticBase(base *lib.BaseSegment, trailer *lib.TrailerInformation) {
	for _, j1 := range base.GetSegments(lib.J1SegmentName) {
		sub := j1.(*lib.J1Segment)
		if sub.ECOACode == lib.ECOACodeZ {
			trailer.TotalECOACodeZ++
		}
		if sub.Validate() == nil {
			trailer.TotalConsumerSegmentsJ1++

			if isValidSocialSecurityNumber(sub.SocialSecurityNumber) {
				trailer.TotalSocialNumbersAllSegments++
				trailer.TotalSocialNumbersJ1Segments++
			}

			if !sub.DateBirth.IsZero() {
				trailer.TotalDatesBirthAllSegments++
				trailer.TotalDatesBirthJ1Segments++
			}

			if sub.TelephoneNumber > 0 {
				trailer.TotalTelephoneNumbersAllSegments++
			}
		}
	}
	for _, j2 := range base.GetSegments(lib.J2SegmentName) {
		sub := j2.(*lib.J2Segment)
		if sub.ECOACode == lib.ECOACodeZ {
			trailer.TotalECOACodeZ++
		}
		if sub.Validate() == nil {
			trailer.TotalConsumerSegmentsJ2++

			if isValidSocialSecurityNumber(sub.SocialSecurityNumber) {
				trailer.TotalSocialNumbersAllSegments++
				trailer.TotalSocialNumbersJ2Segments++
			}

			if !sub.DateBirth.IsZero() {
				trailer.TotalDatesBirthAllSegments++
				trailer.TotalDatesBirthJ2Segments++
			}

			if sub.TelephoneNumber > 0 {
				trailer.TotalTelephoneNumbersAllSegments++
			}
		}
	}
	for _, k1 := range base.GetSegments(lib.K1SegmentName) {
		sub := k1.(*lib.K1Segment)
		if len(sub.OriginalCreditorName) > 0 {
			trailer.TotalOriginalCreditorSegments++
		}
	}
	for _, k2 := range base.GetSegments(lib.K2SegmentName) {
		sub := k2.(*lib.K2Segment)
		if sub.PurchasedIndicator == lib.PurchasedIndicatorToName ||
			sub.PurchasedIndicator == lib.PurchasedIndicatorFromName {
			trailer.TotalPurchasedToSegments++
		}
	}
	for _, k3 := range base.GetSegments(lib.K3SegmentName) {
		sub := k3.(*lib.K3Segment)
		if sub.AgencyIdentifier == lib.AgencyIdentifierNotApplicable {
			trailer.TotalMortgageInformationSegments++
		}
	}
	for _, k4 := range base.GetSegments(lib.K4SegmentName) {
		sub := k4.(*lib.K4Segment)
		if sub.SpecializedPaymentIndicator == lib.SpecializedBalloonPayment ||
			sub.SpecializedPaymentIndicator == lib.SpecializedDeferredPayment {
			trailer.TotalPaymentInformationSegments++
		}
	}
	for _, l1 := range base.GetSegments(lib.L1SegmentName) {
		sub := l1.(*lib.L1Segment)
		if sub.ChangeIndicator == lib.ChangeIndicatorAccountNumber ||
			sub.ChangeIndicator == lib.ChangeIndicatorIdentificationNumber ||
			sub.ChangeIndicator == lib.ChangeIndicatorBothNumber {
			trailer.TotalChangeSegments++
		}
	}
	for _, n1 := range base.GetSegments(lib.N1SegmentName) {
		sub := n1.(*lib.N1Segment)
		if len(sub.EmployerName) > 0 {
			trailer.TotalEmploymentSegments++
		}
	}
}

func isValidSocialSecurityNumber(ssn int) bool {
	// Do not count zero- or 9-filled SSNs.
	if ssn <= 0 || ssn >= 999999999 {
		return false
	}
	return true
}
