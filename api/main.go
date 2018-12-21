package api

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/mailer"
	"github.com/dominik-zeglen/inkster/middleware"
)

type Resolver struct {
	dataSource core.AbstractDataContext
	key        string
	mailer     mailer.Mailer
}

func NewResolver(dataSource core.AbstractDataContext, mailer mailer.Mailer, key string) Resolver {
	return Resolver{
		dataSource: dataSource,
		key:        key,
		mailer:     mailer,
	}
}

type gqlType string

const (
	gqlDirectory gqlType = "directory"
	gqlPage              = "page"
	gqlUser              = "user"
)

func toGlobalID(dataType gqlType, ID int) string {
	data := string(dataType) + ":" + strconv.Itoa(ID)
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func fromGlobalID(dataType gqlType, ID string) (int, error) {
	data, err := base64.StdEncoding.DecodeString(ID)
	if err != nil {
		return 0, err
	}
	portionedData := strings.Split(string(data), ":")
	if portionedData[0] == string(dataType) {
		return strconv.Atoi(portionedData[1])
	}
	return 0, fmt.Errorf("Object types do not match")
}

type userError struct {
	field   string
	message string
}

type userErrorResolver struct {
	data userError
}

func (res *userErrorResolver) Field() string {
	return res.data.field
}

func (res *userErrorResolver) Message() string {
	return res.data.message
}

func checkPermission(ctx context.Context) bool {
	user, ok := ctx.Value("user").(*middleware.UserClaims)
	if ok && user != nil {
		return true
	}
	return false
}

type Sort struct {
	Field string
	Order string
}
