package main

import (
	"bytes"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"testing"
)

var testJsonFilePath = ""

func TestMain(m *testing.M) {
	initRootCmd()
	testJsonFilePath = filepath.Join("../../", "testdata", "packed_file.json")
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

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

func TestConvertJson(t *testing.T) {
	_, err := executeCommand(rootCmd, "convert", "output", "--input", testJsonFilePath, "--format", outputJsonFormat)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestConvertMetro(t *testing.T) {
	_, err := executeCommand(rootCmd, "convert", "output", "--input", testJsonFilePath, "--format", outputMetroFormat)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestConvertMetroWithGenerate(t *testing.T) {
	_, err := executeCommand(rootCmd, "convert", "output", "--input", testJsonFilePath, "--format", outputMetroFormat, "--generate=true")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPrintMetro(t *testing.T) {
	_, err := executeCommand(rootCmd, "print", "--input", testJsonFilePath, "--format", outputMetroFormat)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPrintJson(t *testing.T) {
	_, err := executeCommand(rootCmd, "print", "--input", testJsonFilePath, "--format", outputJsonFormat)
	if err != nil {
		t.Errorf(err.Error())
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
