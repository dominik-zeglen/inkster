package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/urfave/cli.v1"

	"github.com/dominik-zeglen/inkster/api/schema"
	server "github.com/dominik-zeglen/inkster/app"
	"github.com/dominik-zeglen/inkster/core"
)

func main() {
	app := cli.NewApp()
	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			operation := c.Args().Get(0)

			if operation == "runserver" {
				server := server.Server{}
				server.Init(".")

				server.Run()
				return nil
			}

			if operation == "add-user" {
				server := server.Server{}
				server.Init(".")

				dataSource := server.DataSource
				email := c.Args().Get(1)
				password := c.Args().Get(2)
				newUser := core.User{
					Email:  email,
					Active: true,
				}
				newUser.CreatedAt = server.
					DataSource.
					GetCurrentTime()
				newUser.UpdatedAt = server.
					DataSource.
					GetCurrentTime()
				err := newUser.CreatePassword(password)
				if err != nil {
					return err
				}

				_, err = dataSource.
					DB().
					Model(&newUser).
					Insert()

				if err != nil {
					return err
				}
				fmt.Println("Added user " + email)
				return nil
			}

			if operation == "print-schema" {
				fmt.Println(schema.String())
				return nil
			}
		}
		fmt.Println("No operation given")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
