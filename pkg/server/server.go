// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/moov-io/metro2/pkg/file"
	"github.com/moov-io/metro2/pkg/utils"
)

func parseInputFromRequest(r *http.Request) (file.File, error) {
	src, _, err := r.FormFile("file")
	if err != nil {

		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, errors.New("unable to read request body")
		}

		defer r.Body.Close()

		mf, err := file.NewFileFromReader(bytes.NewReader(buf))
		if err != nil {
			return nil, err
		}

		return mf, nil
	}

	defer src.Close()

	mf, err := file.NewFileFromReader(src)
	if err != nil {
		return nil, err
	}
	return mf, nil
}

func outputError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func outputSuccess(w http.ResponseWriter, output string) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": output,
	})
}

func messageToBuf(format string, metroFile file.File, newline bool) ([]byte, error) {
	var output []byte
	var err error
	switch format {
	case utils.MessageJsonFormat:
		output, err = json.MarshalIndent(metroFile, "", "\t")
	case utils.MessageMetroFormat:
		output = []byte(metroFile.String(newline))
	default:
		return nil, errors.New("invalid format")
	}
	return output, err
}

func outputBufferToWriter(w http.ResponseWriter, metroFile file.File, format string, newline bool) {
	w.WriteHeader(http.StatusOK)
	switch format {
	case utils.MessageJsonFormat:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(metroFile)
	case utils.MessageMetroFormat:
		w.Header().Set("Content-Type", "application/octet-stream; charset=utf-8")
		w.Write([]byte(metroFile.String(newline)))
	}
}

func getFormat(r *http.Request) (string, error) {
	var format string
	if r.Header.Get("Content-Type") == "application/json" {
		format = r.URL.Query().Get("format")
	} else {
		format = r.FormValue("format")
	}

	if format == "" {
		format = utils.MessageJsonFormat
	}
	if format != utils.MessageMetroFormat && format != utils.MessageJsonFormat {
		return format, errors.New("invalid format")
	}
	return format, nil
}

func getIsNewLine(r *http.Request) bool {
	newline := true
	newlineStr := r.FormValue("newline")
	if strings.EqualFold(newlineStr, "false") {
		newline = false
	}
	return newline
}

// title: validate metro file
// path: /validator
// method: POST
// produce: multipart/form-data
// responses:
//
//	200: OK
//	400: Bad Request
//	501: Not Implemented
func validator(w http.ResponseWriter, r *http.Request) {
	metroFile, err := parseInputFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = metroFile.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		return
	}

	outputSuccess(w, "valid file")
}

// title: print metro file
// path: /print
// method: POST
// produce: multipart/form-data
// responses:
//
//	200: OK
//	400: Bad Request
//	501: Not Implemented
func print(w http.ResponseWriter, r *http.Request) {
	metroFile, err := parseInputFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	format, err := getFormat(r)
	if err != nil {
		outputError(w, http.StatusNotImplemented, err)
		return
	}

	newline := getIsNewLine(r)
	_, err = messageToBuf(format, metroFile, newline)
	if err != nil {
		outputError(w, http.StatusNotImplemented, err)
		return
	}

	outputBufferToWriter(w, metroFile, format, newline)
}

// title: convert metro file
// path: /convert
// method: POST
// produce: multipart/form-data
// responses:
//
//	200: OK
//	400: Bad Request
//	501: Not Implemented
func convert(w http.ResponseWriter, r *http.Request) {
	metroFile, err := parseInputFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	generate := r.FormValue("generate")
	if strings.EqualFold(generate, "true") {
		trailer, err := metroFile.GeneratorTrailer()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotImplemented)
			return
		}
		err = metroFile.SetRecord(trailer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotImplemented)
			return
		}
	}

	format, err := getFormat(r)
	if err != nil {
		outputError(w, http.StatusNotImplemented, err)
		return
	}

	newType := r.FormValue("type")
	if newType == utils.CharacterFileFormat || newType == utils.PackedFileFormat {
		metroFile.SetType(newType)
	}

	output, err := messageToBuf(format, metroFile, getIsNewLine(r))
	if err != nil {
		outputError(w, http.StatusNotImplemented, err)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=converted_file")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))
}

// title: health server
// path: /health
// method: GET
// responses:
//
//	200: OK
func health(w http.ResponseWriter, r *http.Request) {
	outputSuccess(w, "alive")
}

func ConfigureHandlers() (http.Handler, error) {
	r := mux.NewRouter()
	r.HandleFunc("/health", health).Methods("GET")
	r.HandleFunc("/print", print).Methods("POST")
	r.HandleFunc("/validator", validator).Methods("POST")
	r.HandleFunc("/convert", convert).Methods("POST")
	return r, nil
}
