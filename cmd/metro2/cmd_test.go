// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/bloomcredit/moov-metro2/pkg/utils"
	"github.com/spf13/cobra"
)

var testJsonFilePath = ""

func TestMain(m *testing.M) {
	initRootCmd()
	testJsonFilePath = filepath.Join("..", "..", "test", "testdata", "packed_file.json")
	os.Exit(m.Run())
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}

func deleteFile() {
	// delete file
	os.Remove("output")
}

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

func TestConvertWithoutInput(t *testing.T) {
	_, err := executeCommand(rootCmd, "convert", "output", "--format", utils.MessageJsonFormat)
	if err == nil {
		t.Errorf("invalid input file")
	}
	deleteFile()
}

func TestConvertWithInvalidParam(t *testing.T) {
	_, err := executeCommand(rootCmd, "convert", "--input", testJsonFilePath, "--format", utils.MessageJsonFormat)
	if err == nil {
		t.Errorf("requires output argument")
	}
}

func TestConvertJson(t *testing.T) {
	_, err := executeCommand(rootCmd, "convert", "output", "--input", testJsonFilePath, "--format", utils.MessageJsonFormat)
	if err != nil {
		t.Errorf(err.Error())
	}
	deleteFile()
}

func TestConvertMetro(t *testing.T) {
	_, err := executeCommand(rootCmd, "convert", "output", "--input", testJsonFilePath, "--format", utils.MessageMetroFormat)
	if err != nil {
		t.Errorf(err.Error())
	}
	deleteFile()
}

func TestConvertUnknown(t *testing.T) {
	_, err := executeCommand(rootCmd, "convert", "output", "--input", testJsonFilePath, "--format", "unknown")
	if err == nil {
		t.Errorf("don't support the format")
	}
	deleteFile()
}

func TestConvertMetroWithGenerate(t *testing.T) {
	_, err := executeCommand(rootCmd, "convert", "output", "--input", testJsonFilePath, "--format", utils.MessageMetroFormat, "--generate=true")
	if err != nil {
		t.Errorf(err.Error())
	}
	deleteFile()
}

func TestPrintMetro(t *testing.T) {
	_, err := executeCommand(rootCmd, "print", "--input", testJsonFilePath, "--format", utils.MessageMetroFormat)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPrintJson(t *testing.T) {
	_, err := executeCommand(rootCmd, "print", "--input", testJsonFilePath, "--format", utils.MessageJsonFormat)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPrintUnknown(t *testing.T) {
	_, err := executeCommand(rootCmd, "print", "--input", testJsonFilePath, "--format", "unknown")
	if err == nil {
		t.Errorf("don't support the format")
	}
}

func TestValidator(t *testing.T) {
	_, err := executeCommand(rootCmd, "validator", "--input", testJsonFilePath)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestUnknown(t *testing.T) {
	_, err := executeCommand(rootCmd, "unknown")
	if err == nil {
		t.Errorf("don't support unknown")
	}
}

func TestWeb(t *testing.T) {
	_, err := executeCommand(rootCmd, "web", "--test=true")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestWebWithInvalidPort(t *testing.T) {
	_, err := executeCommand(rootCmd, "web", "--port", "unknown", "--test=false")
	if err == nil {
		t.Errorf("Starting web server with unknown port")
	}
}
