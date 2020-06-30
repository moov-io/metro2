package server

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/moov-io/metro2/file"
	"github.com/moov-io/metro2/utils"
)

func parseInput(c echo.Context) (file.File, error) {
	f, err := c.FormFile("file")
	if err != nil {
		return nil, c.JSON(http.StatusBadRequest, err.Error())
	}
	src, err := f.Open()
	if err != nil {
		return nil, c.JSON(http.StatusBadRequest, err.Error())
	}
	defer src.Close()

	var input bytes.Buffer
	if _, err = io.Copy(&input, src); err != nil {
		return nil, c.JSON(http.StatusBadRequest, err.Error())
	}

	space := regexp.MustCompile(`\s+`)
	buf := space.ReplaceAllString(input.String(), " ")
	mf, err := file.CreateFile([]byte(buf))
	if err != nil {
		return nil, c.JSON(http.StatusNotImplemented, err.Error())
	}
	return mf, nil
}

// title: validate metro file
// path: /validator
// method: POST
// produce: multipart/form-data
// responses:
//   200: OK
//   400: Bad Request
//   501: Not Implemented
func validator(c echo.Context) error {
	mf, err := parseInput(c)
	if err != nil {
		return err
	}

	err = mf.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "valid file")
}

// title: print metro file
// path: /print
// method: POST
// produce: multipart/form-data
// responses:
//   200: OK
//   400: Bad Request
//   501: Not Implemented
func print(c echo.Context) error {
	mf, err := parseInput(c)
	if err != nil {
		return err
	}

	format := c.FormValue("format")
	if format == utils.OutputMetroFormat {
		return c.JSON(http.StatusOK, mf)
	}

	return c.JSONPretty(http.StatusOK, mf, "  ")
}

// title: convert metro file
// path: /convert
// method: POST
// produce: multipart/form-data
// responses:
//   200: OK
//   400: Bad Request
//   501: Not Implemented
func convert(c echo.Context) error {
	mf, err := parseInput(c)
	if err != nil {
		return err
	}

	generate := c.FormValue("generate")
	if generate == "true" {
		trailer, err := mf.GeneratorTrailer()
		if err != nil {
			return c.JSON(http.StatusNotImplemented, err.Error())
		}
		err = mf.SetRecord(trailer)
		if err != nil {
			return c.JSON(http.StatusNotImplemented, err.Error())
		}
	}

	format := c.FormValue("format")
	buf, err := json.Marshal(mf)
	if err != nil {
		return c.JSON(http.StatusNotImplemented, err.Error())
	}

	filename := "metro.json"
	output := string(buf)
	if format == utils.OutputMetroFormat {
		output = mf.String()
		filename = "metro"
	}

	res := c.Response()
	writer := bufio.NewWriter(res)
	header := res.Header()
	header.Set(echo.HeaderContentType, echo.MIMEOctetStream)
	header.Set(echo.HeaderContentDisposition, "attachment; filename="+filename)
	header.Set("Content-Transfer-Encoding", "binary")
	header.Set("Expires", "0")
	res.WriteHeader(http.StatusOK)

	writer.WriteString(output)
	writer.Flush()

	return nil
}

// title: health server
// path: /health
// method: GET
// responses:
//   200: OK
func health(c echo.Context) error {
	return c.JSON(http.StatusOK, "health")
}

func ConfigureHandlers() (http.Handler, error) {
	r := echo.New()
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	r.POST("/validator", validator)
	r.POST("/print", print)
	r.POST("/convert", convert)
	r.GET("/health", health)

	return r, nil
}
