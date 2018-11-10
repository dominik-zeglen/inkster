package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-testfixtures/testfixtures"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_HOST"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Print("Generating fixtures from database")

	err = testfixtures.GenerateFixtures(
		db,
		&testfixtures.PostgreSQL{},
		"fixtures",
	)
	if err != nil {
		log.Fatalf("Error generating fixtures: %v", err)
	} else {
		log.Printf("Fixtures generated successfully")
	}
}
