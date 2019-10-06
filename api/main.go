package api

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/mail"
	"github.com/dominik-zeglen/inkster/middleware"
)

type Resolver struct {
	dataSource core.AbstractDataContext
	key        string
	mailer     mail.Mailer
}

func NewResolver(dataSource core.AbstractDataContext, mailer mail.Mailer, key string) Resolver {
	return Resolver{
		dataSource,
		key,
		mailer,
	}
}

type gqlType string

const (
	gqlDirectory gqlType = "directory"
	gqlPage              = "page"
	gqlPageField         = "pageField"
	gqlUser              = "user"
	gqlCursor            = "cursor"
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

func toGlobalCursor(ID Cursor) string {
	return toGlobalID(gqlCursor, int(ID))
}

func fromGlobalCursor(cursor string) (Cursor, error) {
	data, err := fromGlobalID(gqlCursor, cursor)

	return Cursor(data), err
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
	user, ok := ctx.Value(middleware.UserContextKey).(*core.User)
	if ok && user != nil {
		return true
	}
	return false
}

type Cursor int32

type Sort struct {
	Field string
	Order string
}

type Paginate struct {
	After  *Cursor
	Before *Cursor
	First  *int32
	Last   *int32
}

type PageInfo struct {
	endCursor       *Cursor
	hasNextPage     bool
	hasPreviousPage bool
	startCursor     *Cursor
}
