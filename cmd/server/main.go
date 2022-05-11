package main

import (
	"fmt"
	"log"
)

// Run - is responsible for the instantiation and startup
// of our go application
func Run() error {
	fmt.Println("starting our application")
	return nil
}

func main() {
	fmt.Println("Go REST API")
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
