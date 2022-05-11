package main

import (
	"context"
	"fmt"
	"log"

	"github.com/snirkop89/go-restapi/internal/comment"
	"github.com/snirkop89/go-restapi/internal/db"
)

// Run - is responsible for the instantiation and startup
// of our go application
func Run() error {
	fmt.Println("starting our application")

	// give a few seconds for the postgres docker to start
	db, err := db.NewDatabase()
	if err != nil {
		log.Println("Failed to connect to the database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		log.Println("failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)
	log.Println(cmtService.GetComment(context.Background(), "71c5d074-b6cf-11ec-b909-0242ac120002"))

	return nil
}

func main() {
	fmt.Println("Go REST API")
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
