package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", pong)
	http.HandleFunc("/hello", helloWorld)

	log.Println("Starting http server ...")
	log.Fatal(http.ListenAndServe(":50052", nil))

}

func pong(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("pong\n"))
	if err != nil {
		return
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello world\n"))
	if err != nil {
		return
	}
}
