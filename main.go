package main

import (
	"fmt"
	"net/http"
)

func pathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.path=%s\nURL.RawPath=%s", r.URL.Path, r.URL.RawPath)
}

func main() {
	http.HandleFunc("/", pathHandler)
	fmt.Println("Starting the server on :3000...")
	_ = http.ListenAndServe(":3000", nil)
}
