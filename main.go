package main

import (
	"net/http"
	"log"
	"fmt"
	"time"
	"github.com/tj/go-spin"
)

const Addr = ":8080"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(createMessage())
	})

	fmt.Println("Starting app on address " + Addr)
	go startSpinner()

	log.Fatal(http.ListenAndServe(Addr, nil))
}

func createMessage() []byte {
	return []byte("Hello!")
}

func startSpinner() {
	s := spin.New()
	for {
		fmt.Printf("\r  \033[33mServing Traffic!\033[m %s ", s.Next())
		time.Sleep(100 * time.Millisecond)
	}
}
