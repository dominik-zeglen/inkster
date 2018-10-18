package api

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func closeQuietly(v interface{}) {
	if d, ok := v.(io.Closer); ok {
		_ = d.Close()
	}
}
func TestUpload(t *testing.T) {
	t.Run("Test valid request", func(t *testing.T) {
		file, err := os.Open("testdata/testfile.test")
		if err != nil {
			t.Fatal(err)
		}
		closeQuietly(file)

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("file", filepath.Base("testdata/testfile.test"))
		if err != nil {
			t.Fatal(err)
		}
		_, err = io.Copy(part, file)
		if err != nil {
			t.Fatal(err)
		}
		err = writer.Close()
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/upload", body)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-type", writer.FormDataContentType())
		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(UploadHandler)
		handler.ServeHTTP(responseRecorder, req)

		status := responseRecorder.Code
		if status != http.StatusOK {
			t.Errorf("Handler output: %v", responseRecorder.Body)
			t.Fatalf(
				"Handler returned wrong status code: got %v, want %v",
				status,
				http.StatusOK,
			)
		}
	})
	t.Run("Test GET request", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/upload", nil)
		if err != nil {
			t.Fatal(err)
		}

		responseRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(UploadHandler)
		handler.ServeHTTP(responseRecorder, req)

		status := responseRecorder.Code
		if status != http.StatusBadRequest {
			t.Errorf(
				"Handler returned wrong status code: got %v, want %v",
				status,
				http.StatusBadRequest,
			)
		}
	})
}
