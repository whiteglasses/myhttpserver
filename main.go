package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	err := http.ListenAndServe(":3333", mux)
	if !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
