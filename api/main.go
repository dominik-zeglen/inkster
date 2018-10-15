package api

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/mailer"
	"github.com/dominik-zeglen/inkster/middleware"
	"github.com/globalsign/mgo/bson"
)

type Resolver struct {
	dataSource core.Adapter
	key        string
	mailer     mailer.Mailer
}

func NewResolver(dataSource core.Adapter, mailer mailer.Mailer, key string) Resolver {
	return Resolver{
		dataSource: dataSource,
		key:        key,
		mailer:     mailer,
	}
}

func toGlobalID(dataType string, ID bson.ObjectId) string {
	data := dataType + ":" + string(ID)
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func fromGlobalID(dataType string, ID string) (bson.ObjectId, error) {
	data, err := base64.StdEncoding.DecodeString(ID)
	if err != nil {
		panic(err)
	}
	portionedData := strings.Split(string(data), ":")
	if portionedData[0] == dataType {
		return bson.ObjectId(portionedData[1]), nil
	}
	return "", fmt.Errorf("Object types do not match")
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
