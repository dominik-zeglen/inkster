package api

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

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

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(400)
		_, err := w.Write(net.NewNetworkError(http.ErrNoLocation).ToJson())
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	err := r.ParseMultipartForm((1 << 20) * 10) // 10MB max size
	if sendError(err, w, 400) {
		return
	}
	file, fileHeader, err := r.FormFile("file")
	if sendError(err, w, 500) {
		return
	}
	defer closeFile(file, w)

	filename, err := createFileName(file)
	if sendError(err, w, 500) {
		return
	}
	f, err := os.OpenFile(
		os.Getenv("INKSTER_STATIC")+"/"+filename+"_"+fileHeader.Filename,
		os.O_WRONLY|os.O_CREATE,
		0666,
	)
	if sendError(err, w, 500) {
		return
	}
	defer closeFile(f, w)
	_, err = io.Copy(f, file)
	if sendError(err, w, 500) {
		return
	}
	_, err = w.Write([]byte("{\"filename\":\"" + filename + "_" + fileHeader.Filename + "\"}"))
	if err != nil {
		log.Fatal(err)
	}
	return
}
