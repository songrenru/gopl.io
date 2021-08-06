// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", 哈哈)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func 哈哈(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}