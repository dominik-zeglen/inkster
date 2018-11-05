package api

import "github.com/dominik-zeglen/inkster/core"

type inputErrorResolver struct {
	err core.ValidationError
}

func (res inputErrorResolver) Code() int32 {
	return int32(res.err.Code)
}

func (res inputErrorResolver) Field() string {
	return res.err.Field
}

func (res inputErrorResolver) Message() string {
	return res.err.Error()
}
