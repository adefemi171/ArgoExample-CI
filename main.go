package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test CI Webserver 3")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":3000", nil)
}
