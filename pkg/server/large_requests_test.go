package server

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestServer__LargeRequests(t *testing.T) {
	timeout, _ := time.ParseDuration("30s")
	handler, _ := ConfigureHandlers()

	svr := &http.Server{
		Addr:              "0.0.0.0:15551",
		Handler:           handler,
		ReadTimeout:       timeout,
		ReadHeaderTimeout: timeout,
		WriteTimeout:      timeout,
		IdleTimeout:       timeout,
	}
	go svr.ListenAndServe()
	t.Cleanup(func() {
		svr.Shutdown(context.Background())
	})

	// Read each file
	files := []string{
		filepath.Join("testdata", "10k_record.json.gz"),
		filepath.Join("testdata", "25k_record.json.gz"),
		filepath.Join("testdata", "50k_record.json.gz"),
	}
	for i := range files {
		path := files[i]

		t.Run("validate "+path, func(t *testing.T) {
			var body bytes.Buffer
			w := multipart.NewWriter(&body)
			part, err := w.CreateFormFile("file", filepath.Base(path))
			require.NoError(t, err)
			_, err = io.Copy(part, open(t, path))
			require.NoError(t, err)
			require.NoError(t, w.Close()) // flush

			req, err := http.NewRequest("POST", "http://localhost:15551/validator", &body)
			require.NoError(t, err)
			req.Header.Set("Content-Type", w.FormDataContentType())

			resp, err := http.DefaultClient.Do(req)
			if resp != nil && resp.StatusCode != http.StatusOK {
				if resp != nil && resp.Body != nil {
					t.Cleanup(func() { resp.Body.Close() })

					bs, _ := io.ReadAll(resp.Body)
					t.Logf("Response: %v", string(bs))
				}
				require.Equal(t, http.StatusOK, resp.StatusCode)
			}
			require.NoError(t, err)
		})

		t.Run("convert "+path, func(t *testing.T) {
			var body bytes.Buffer
			w := multipart.NewWriter(&body)
			part, err := w.CreateFormFile("file", filepath.Base(path))
			require.NoError(t, err)
			_, err = io.Copy(part, open(t, path))
			require.NoError(t, err)
			w.WriteField("format", "metro")
			require.NoError(t, w.Close()) // flush

			req, err := http.NewRequest("POST", "http://localhost:15551/convert", &body)
			require.NoError(t, err)
			req.Header.Set("Content-Type", w.FormDataContentType())

			resp, err := http.DefaultClient.Do(req)
			if resp != nil && resp.StatusCode != http.StatusOK {
				if resp != nil && resp.Body != nil {
					t.Cleanup(func() { resp.Body.Close() })

					bs, _ := io.ReadAll(resp.Body)
					t.Logf("Response: %v", string(bs))
				}
				require.Equal(t, http.StatusOK, resp.StatusCode)
			}
			require.NoError(t, err)
		})
	}
}

func open(t *testing.T, path string) io.Reader {
	t.Helper()

	fd, err := os.Open(path)
	require.NoError(t, err)
	t.Cleanup(func() { fd.Close() })

	r, err := gzip.NewReader(fd)
	require.NoError(t, err)

	return r
}
