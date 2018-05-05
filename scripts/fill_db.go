package util

import (
	"fmt"
	"os"

	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/dominik-zeglen/ecoknow/postgres"
	"github.com/go-pg/pg"
)

func FillDB() {
	dataSource := postgres.Adapter{ConnectionOptions: pg.Options{
		User:     os.Getenv("FOXXY_DB_USER"),
		Password: os.Getenv("FOXXY_DB_USER_PASSWORD"),
		Database: os.Getenv("FOXXY_DB_NAME"),
		Addr: fmt.Sprintf("%s:%s",
			os.Getenv("FOXXY_DB_ADDR"),
			os.Getenv("FOXXY_DB_PORT"),
		),
	}}
	containerNames := []string{"Lorem", "Ipsum", "Dolor", "Sit", "Amet"}
	for _, name := range containerNames {
		container := core.Container{Name: name}
		container, err := dataSource.AddContainer(container)
		if err != nil {
			panic(err)
		}
	}

	containerChildrenNames := []string{"Lorem Lorem", "Lorem Ipsum", "Lorem Dolor", "Lorem Sit", "Lorem Amet"}
	for _, name := range containerChildrenNames {
		container := core.Container{Name: name, ParentID: 1}
		container, err := dataSource.AddContainer(container)
		if err != nil {
			panic(err)
		}
	}
}
