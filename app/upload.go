package app

import (
	"net/http"
	"time"

	"github.com/dominik-zeglen/inkster/api"
	"github.com/dominik-zeglen/inkster/storage"
)

type UploadHandler struct {
	http.Handler
	fileUploader   storage.FileUploader
	getCurrentTime func() time.Time
}

func (handler UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.UploadHandler(w, r, handler.fileUploader, handler.getCurrentTime().String())
}

func newUploadHandler(
	fileUploader storage.FileUploader,
	getCurrentTime func() time.Time,
) UploadHandler {
	return UploadHandler{
		fileUploader:   fileUploader,
		getCurrentTime: getCurrentTime,
	}
}
