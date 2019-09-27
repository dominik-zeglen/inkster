package app

import (
	"net/http"

	"github.com/dominik-zeglen/inkster/api"
	"github.com/dominik-zeglen/inkster/storage"
)

type UploadHandler struct {
	http.Handler
	fileUploader storage.FileUploader
}

func (handler UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.UploadHandler(w, r, handler.fileUploader)
}

func newUploadHandler(
	fileUploader storage.FileUploader,
) UploadHandler {
	return UploadHandler{
		fileUploader: fileUploader,
	}
}
