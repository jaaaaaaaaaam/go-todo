package main

import (
	"github.com/jaaaaaaaaaam/go-todo/router"
	"log"
	"net/http"
)

func main() {

	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

}
