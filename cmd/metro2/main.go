// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/moov-io/metro2/pkg/file"
	"github.com/moov-io/metro2/pkg/server"
	"github.com/moov-io/metro2/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	inputFile = ""
	rawData   = ""
)

var WebCmd = &cobra.Command{
	Use:   "web",
	Short: "Launches web server",
	Long:  "Launches web server",
	RunE: func(cmd *cobra.Command, args []string) error {
		port, err := cmd.Flags().GetString("port")
		if err != nil {
			return err
		}
		fmt.Println("Starting web server on port ", port)
		listen := "0.0.0.0:" + port
		h, err := server.ConfigureHandlers()
		if err != nil {
			return err
		}
		test, _ := cmd.Flags().GetBool("test")
		if !test {
			log.Println(http.ListenAndServe(listen, h))
		}
		return nil
	},
}

var Validate = &cobra.Command{
	Use:   "validator",
	Short: "Validate metro file",
	Long:  "Validate an incoming metro file",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := file.CreateFile([]byte(rawData))
		if err != nil {
			return err
		}
		return f.Validate()
	},
}

var Print = &cobra.Command{
	Use:   "print",
	Short: "Print metro file",
	Long:  "Print an incoming metro file with special format (options: metro, json)",
	RunE: func(cmd *cobra.Command, args []string) error {
		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return err
		}
		if format != utils.OutputJsonFormat && format != utils.OutputMetroFormat {
			return errors.New("don't support the format")
		}

		f, err := file.CreateFile([]byte(rawData))
		if err != nil {
			return err
		}

		output := ""
		if format == utils.OutputJsonFormat {
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
		} else if format == utils.OutputMetroFormat {
			output = f.String()
		}
		fmt.Println(output)
		return nil
	},
}

var Convert = &cobra.Command{
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
		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return err
		}
		if format != utils.OutputJsonFormat && format != utils.OutputMetroFormat {
			return errors.New("don't support the format")
		}

		mf, err := file.CreateFile([]byte(rawData))
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

		output := ""
		if format == utils.OutputJsonFormat {
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
		} else if format == utils.OutputMetroFormat {
			output = mf.String()
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

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		isWeb := false
		cmdNames := make([]string, 0)
		getName := func(c *cobra.Command) {}
		getName = func(c *cobra.Command) {
			if c == nil {
				return
			}
			cmdNames = append([]string{c.Name()}, cmdNames...)
			if c.Name() == "web" {
				isWeb = true
			}
			getName(c.Parent())
		}
		getName(cmd)

		if !isWeb {
			if inputFile == "" {
				path, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}
				inputFile = filepath.Join(path, "metro.json")
			}
			_, err := os.Stat(inputFile)
			if os.IsNotExist(err) {
				return errors.New("invalid input file")
			}
			f, err := os.Open(inputFile)
			if err != nil {
				return err
			}
			rawData = utils.ReadFile(f)
		}

		return nil
	},
}

func initRootCmd() {
	WebCmd.Flags().String("port", "8080", "port of the web server")
	WebCmd.Flags().BoolP("test", "t", false, "test server")
	Convert.Flags().String("format", "json", "format of metro file(required)")
	Convert.Flags().BoolP("generate", "g", false, "generate trailer record")
	Convert.MarkFlagRequired("format")
	Print.Flags().String("format", "json", "print format")

	rootCmd.SilenceUsage = true
	rootCmd.PersistentFlags().StringVar(&inputFile, "input", "", "input file (default is $PWD/metro.json)")
	rootCmd.AddCommand(WebCmd)
	rootCmd.AddCommand(Convert)
	rootCmd.AddCommand(Print)
	rootCmd.AddCommand(Validate)
}

func main() {
	initRootCmd()

	rootCmd.Execute()
}
