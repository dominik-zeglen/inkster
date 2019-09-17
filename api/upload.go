package api

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"

	"github.com/dominik-zeglen/inkster/net"
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

func createFileName(file multipart.File) (string, error) {
	hash := md5.New()
	_, err := io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
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
	filename string
}

// UploadHandler is handler that allows dropping files to Inkster
func UploadHandler(w http.ResponseWriter, r *http.Request) {
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
	var buff bytes.Buffer
	_, err = buff.ReadFrom(file)
	if sendError(err, w, 500) {
		return
	}

	filename, err := createFileName(file)
	savedFilename := filename + "_" + fileHeader.Filename
	if sendError(err, w, 500) {
		return
	}
	f, err := os.OpenFile(
		path.Join("../static", savedFilename),
		os.O_WRONLY|os.O_CREATE,
		0666,
	)
	if sendError(err, w, 500) {
		return
	}
	defer closeFile(f, w)

	_, err = f.Write(buff.Bytes())
	if sendError(err, w, 500) {
		return
	}

	res, err := json.Marshal(UploadResponse{
		filename: savedFilename,
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write([]byte(res))
	if err != nil {
		log.Fatal(err)
	}
	return
}
