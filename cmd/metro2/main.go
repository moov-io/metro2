// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"

	"github.com/moov-io/metro2/pkg/file"
	"github.com/moov-io/metro2/pkg/server"
	"github.com/moov-io/metro2/pkg/utils"
)

// utility functions
func createReader(cmd *cobra.Command) (io.Reader, error) {

	input, err := cmd.Flags().GetString("input")
	if err != nil {
		return nil, err
	}

	if input == "" {
		path, _ := os.Getwd()
		input = filepath.Join(path, "metro.json")
	}

	_, err = os.Stat(input)
	if os.IsNotExist(err) {
		return nil, errors.New("invalid input file")
	}
	reader, err := os.Open(input)
	if err != nil {
		return nil, err
	}

	return reader, nil
}

func closeReader(reader io.Reader) {
	if closer, ok := reader.(io.Closer); ok {
		closer.Close()
	}
}

// commands
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Launches web server",
	Long:  "Launches web server",
	RunE: func(cmd *cobra.Command, args []string) error {

		port, _ := cmd.Flags().GetString("port")
		fmt.Fprintf(os.Stdout, "Starting web server on port %s\n\n", port)

		timeout, _ := time.ParseDuration("30s")
		handler, _ := server.ConfigureHandlers()
		serve := &http.Server{
			Addr:              "0.0.0.0:" + port,
			Handler:           handler,
			ReadTimeout:       timeout,
			ReadHeaderTimeout: timeout,
			WriteTimeout:      timeout,
			IdleTimeout:       timeout,
		}

		test, _ := cmd.Flags().GetBool("test")
		if !test {
			if err := serve.ListenAndServe(); err != nil {
				return err
			}
		}
		return nil
	},
}

var validate = &cobra.Command{
	Use:   "validator",
	Short: "Validate metro file",
	Long:  "Validate an incoming metro file",
	RunE: func(cmd *cobra.Command, args []string) error {

		reader, err := createReader(cmd)
		if err != nil {
			return err
		}
		defer closeReader(reader)

		f, err := file.NewFileFromReader(reader)
		if err != nil {
			return err
		}

		err = f.Validate()
		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "the file is valid \n")

		return nil
	},
}

var print = &cobra.Command{
	Use:   "print",
	Short: "Print metro file",
	Long:  "Print an incoming metro file with special format (options: metro, json)",
	RunE: func(cmd *cobra.Command, args []string) error {

		reader, err := createReader(cmd)
		if err != nil {
			return err
		}
		defer closeReader(reader)

		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return err
		}

		if format != utils.MessageJsonFormat && format != utils.MessageMetroFormat {
			if format == "" {
				format = utils.MessageJsonFormat
			} else {
				return errors.New("don't support the format")
			}
		}

		f, err := file.NewFileFromReader(reader)
		if err != nil {
			return err
		}

		newline, _ := cmd.Flags().GetBool("newline")
		output := ""

		if format == utils.MessageJsonFormat {
			buf, err := json.Marshal(f)
			if err != nil {
				return err
			}
			var pretty bytes.Buffer
			err = json.Indent(&pretty, buf, "", "  ")
			if err != nil {
				return err
			}
			output = pretty.String()
		} else if format == utils.MessageMetroFormat {
			output = f.String(newline)
		}

		fmt.Fprintf(os.Stdout, "%s", output)

		return nil
	},
}

var convert = &cobra.Command{
	Use:   "convert [output]",
	Short: "Convert metro file format",
	Long:  "Convert an incoming metro file into another format (options: metro, json)",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires output argument")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		reader, err := createReader(cmd)
		if err != nil {
			return err
		}
		defer closeReader(reader)

		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return err
		}

		if format != utils.MessageJsonFormat && format != utils.MessageMetroFormat {
			if format == "" {
				format = utils.MessageJsonFormat
			} else {
				return errors.New("don't support the format")
			}
		}

		mf, err := file.NewFileFromReader(reader)
		if err != nil {
			return err
		}

		generate, _ := cmd.Flags().GetBool("generate")
		if generate {
			trailer, err := mf.GeneratorTrailer()
			if err != nil {
				return err
			}
			err = mf.SetRecord(trailer)
			if err != nil {
				return err
			}
		}

		newline, _ := cmd.Flags().GetBool("newline")
		if newType, tErr := cmd.Flags().GetString("type"); tErr == nil {
			if newType == utils.PackedFileFormat || newType == utils.CharacterFileFormat {
				mf.SetType(newType)
			}
		}

		output := ""
		if format == utils.MessageJsonFormat {
			buf, err := json.Marshal(mf)
			if err != nil {
				return err
			}
			var pretty bytes.Buffer
			err = json.Indent(&pretty, buf, "", "  ")
			if err != nil {
				return err
			}
			output = pretty.String()
		} else if format == utils.MessageMetroFormat {
			output = mf.String(newline)
		}

		f, err := os.Create(args[0])
		if err != nil {
			return err
		}
		_, err = f.WriteString(output)
		f.Close()
		return err
	},
}

var rootCmd = &cobra.Command{}

func initRootCmd() {
	webCmd.Flags().String("port", "8080", "port of the web server")
	webCmd.Flags().BoolP("test", "t", false, "test server")

	convert.Flags().String("format", "json", "format of metro file(required)")
	convert.Flags().String("type", "", "file type (character or packed)")
	convert.Flags().BoolP("generate", "g", false, "generate trailer record")
	convert.Flags().BoolP("newline", "n", false, "has new line")

	print.Flags().String("format", "json", "print format")
	print.Flags().BoolP("newline", "n", false, "has new line")

	rootCmd.SilenceUsage = true
	rootCmd.PersistentFlags().String("input", "", "input file (default is $PWD/metro.json)")

	rootCmd.AddCommand(webCmd)
	rootCmd.AddCommand(convert)
	rootCmd.AddCommand(print)
	rootCmd.AddCommand(validate)
}

func main() {
	initRootCmd()

	rootCmd.Execute()
}
