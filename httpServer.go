package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// register handle functions
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)

	// listen and serve on specified PORT
	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello gopher on the web")
}

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}
