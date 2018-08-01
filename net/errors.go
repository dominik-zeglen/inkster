package net

import (
	"encoding/json"
	"net/http"
)

type NetworkError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewNetworkError(err error) NetworkError {
	e := NetworkError{
		Message: err.Error(),
	}
	if err == http.ErrMissingFile {
		e.Code = "ErrMissingFile"
	}
	if e.Code == "" {
		e.Code = "ErrUnknown"
	}
	return e
}
func (uploadError NetworkError) ToJson() []byte {
	out, err := json.Marshal(&uploadError)
	if err != nil {
		panic(err)
	}
	return out
}
