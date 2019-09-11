package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
