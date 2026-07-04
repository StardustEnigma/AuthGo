package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/StardustEnigma/AuthGo/db"
	"github.com/StardustEnigma/AuthGo/routes"
)

func main() {
	fmt.Println("starting server at port :8080")
	err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	routes.Routes()
	if err := http.ListenAndServe(":8080",nil); err != nil{
		log.Fatal(err)
	}
}