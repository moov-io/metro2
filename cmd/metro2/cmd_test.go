// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/metro2"
	"github.com/moov-io/metro2/pkg/utils"
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
		t.Error(err)
	}
	deleteFile()
}

func TestConvertMetro(t *testing.T) {
	_, err := executeCommand(rootCmd, "convert", "output", "--input", testJsonFilePath, "--format", utils.MessageMetroFormat)
	if err != nil {
		t.Error(err)
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
		t.Error(err)
	}
	deleteFile()
}

func TestPrintMetro(t *testing.T) {
	_, err := executeCommand(rootCmd, "print", "--input", testJsonFilePath, "--format", utils.MessageMetroFormat)
	if err != nil {
		t.Error(err)
	}
}

func TestPrintJson(t *testing.T) {
	_, err := executeCommand(rootCmd, "print", "--input", testJsonFilePath, "--format", utils.MessageJsonFormat)
	if err != nil {
		t.Error(err)
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
		t.Error(err)
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
		t.Error(err)
	}
}

func TestWebWithInvalidPort(t *testing.T) {
	_, err := executeCommand(rootCmd, "web", "--port", "unknown", "--test=false")
	if err == nil {
		t.Errorf("Starting web server with unknown port")
	}
}

func TestVersion(t *testing.T) {
	output, err := executeCommand(rootCmd, "version")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(output, metro2.Version) {
		t.Errorf("expected version %q in output, got %q", metro2.Version, output)
	}
}

func TestHelpProgramName(t *testing.T) {
	output, err := executeCommand(rootCmd, "--help")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(output, "metro2 [command]") {
		t.Errorf("expected program name in usage, got %q", output)
	}
}

func TestCLIExitCodes(t *testing.T) {
	bin := filepath.Join(t.TempDir(), "metro2")
	build := exec.Command("go", "build", "-o", bin, ".")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}

	testdata := filepath.Join("..", "..", "test", "testdata", "packed_file.json")

	t.Run("validator missing input", func(t *testing.T) {
		cmd := exec.Command(bin, "validator", "--input", "nope.dat")
		out, err := cmd.CombinedOutput()
		if err == nil {
			t.Fatalf("expected non-zero exit, got 0, output %s", out)
		}
		var ee *exec.ExitError
		if !errors.As(err, &ee) || ee.ExitCode() != 1 {
			t.Fatalf("expected exit 1, got err=%v out=%s", err, out)
		}
		if !strings.Contains(string(out), "invalid input file") {
			t.Fatalf("unexpected output %s", out)
		}
	})

	t.Run("convert missing directory", func(t *testing.T) {
		outPath := filepath.Join(t.TempDir(), "missing", "output.json")
		cmd := exec.Command(bin, "convert", outPath, "--input", testdata, "--format", "json")
		out, err := cmd.CombinedOutput()
		if err == nil {
			t.Fatalf("expected non-zero exit, got 0, output %s", out)
		}
		var ee *exec.ExitError
		if !errors.As(err, &ee) || ee.ExitCode() != 1 {
			t.Fatalf("expected exit 1, got err=%v out=%s", err, out)
		}
	})

	t.Run("validator success", func(t *testing.T) {
		cmd := exec.Command(bin, "validator", "--input", testdata)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("expected exit 0, got %v out=%s", err, out)
		}
		if !strings.Contains(string(out), "the file is valid") {
			t.Fatalf("unexpected output %s", out)
		}
	})
}
