package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/snirkop89/go-restapi/internal/db"
)

// Run - is responsible for the instantiation and startup
// of our go application
func Run() error {
	fmt.Println("starting our application")

	var dbConn *db.Database
	var err error
	// give a few seconds for the postgres docker to start
	for i := 0; i < 2; i++ {
		dbConn, err = db.NewDatabase()
		if err != nil {
			log.Println("Failed to connect to the database")
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return err
	}
	if err := dbConn.Ping(context.Background()); err != nil {
		log.Println(err)
	}

	return nil
}

func main() {
	fmt.Println("Go REST API")
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
