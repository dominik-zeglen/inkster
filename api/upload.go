package api

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/dominik-zeglen/inkster/net"
	"github.com/dominik-zeglen/inkster/storage"
)

func closeFile(file interface{ Close() error }, w http.ResponseWriter) {
	err := file.Close()
	if err != nil {
		w.WriteHeader(500)
		_, err := w.Write(net.NewNetworkError(err).ToJson())
		if err != nil {
			log.Fatal(err)
		}
	}
}

func createFileName(filename string, datetime string) (string, error) {
	return fmt.Sprintf(
		"%x_%x%s",
		md5.Sum([]byte(filename)),
		md5.Sum([]byte(datetime)),
		filepath.Ext(filename),
	), nil
}

func sendError(err error, w http.ResponseWriter, code int) bool {
	if err != nil {
		w.WriteHeader(code)
		_, err := w.Write(net.NewNetworkError(err).ToJson())
		if err != nil {
			log.Fatal(err)
		}
		return true
	}
	return false
}

// UploadResponse is a type of success response of UploadHandler
type UploadResponse struct {
	Filename string `json:"filename"`
}

// UploadHandler is handler that allows dropping files to Inkster
func UploadHandler(
	w http.ResponseWriter,
	r *http.Request,
	uploader storage.FileUploader,
	currentTime string,
) {
	if r.Method == "GET" {
		w.WriteHeader(400)
		_, err := w.Write(net.NewNetworkError(http.ErrNoLocation).ToJson())
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	err := r.ParseMultipartForm(32 << 20)
	if sendError(err, w, 400) {
		return
	}
	file, fileHeader, err := r.FormFile("file")
	if sendError(err, w, 500) {
		return
	}
	defer closeFile(file, w)

	filename, err := createFileName(fileHeader.Filename, currentTime)
	if sendError(err, w, 500) {
		return
	}

	url, err := uploader.Upload(file, filename)
	if sendError(err, w, 500) {
		return
	}

	res, err := json.Marshal(UploadResponse{
		Filename: url,
	})
	if sendError(err, w, 500) {
		return
	}

	_, err = w.Write([]byte(res))
	if sendError(err, w, 500) {
		return
	}

	return
}
