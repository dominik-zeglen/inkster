package storage

import "io"

type FileUploader interface {
	Upload(io.Reader, string) (string, error)
}
