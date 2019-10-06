package main

import (
	"./models"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %s!", r.URL.Path[1:])
}

func main() {
	fmt.Sprint(models.GetDB());
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

