package main

import (
	"net/http"
	"log"
	"fmt"
)

const Addr = ":8080"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(createMessage())
	})

	fmt.Println("Starting app on address " + Addr)
	log.Fatal(http.ListenAndServe(Addr, nil))
}

func createMessage() []byte {
	return []byte("Hello!")
}
