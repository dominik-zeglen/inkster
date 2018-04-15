package core

import (
	"fmt"
	"os"

	"github.com/go-pg/pg"
)

var DbOptions = &pg.Options{
	User:     os.Getenv("FOXXY_DB_USER"),
	Password: os.Getenv("FOXXY_DB_USER_PASSWORD"),
	Database: os.Getenv("FOXXY_DB_NAME"),
	Addr: fmt.Sprintf("%s:%s",
		os.Getenv("FOXXY_DB_ADDR"),
		os.Getenv("FOXXY_DB_PORT"),
	),
}
