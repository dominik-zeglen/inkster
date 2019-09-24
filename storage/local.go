package storage

import (
	"bytes"
	"io"
	"os"
	"path"
)

type LocalFileUploader struct{}

func NewLocalFileUploader() LocalFileUploader {
	return LocalFileUploader{}
}

func (uploader LocalFileUploader) Upload(
	file io.Reader,
	filename string,
) (string, error) {
	f, err := os.OpenFile(
		path.Join("static", filename),
		os.O_WRONLY|os.O_CREATE,
		0666,
	)
	if err != nil {
		return filename, err
	}

	defer f.Close()

	var buff bytes.Buffer
	_, err = buff.ReadFrom(file)
	if err != nil {
		return filename, err
	}

	_, err = f.Write(buff.Bytes())
	if err != nil {
		return filename, err
	}

	return path.Join("/static", filename), nil
}
