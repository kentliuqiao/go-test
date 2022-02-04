package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{store: NewInMemoryStore()}
	log.Fatal(http.ListenAndServe(":5001", server))
}
