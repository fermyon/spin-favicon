package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	spin "github.com/fermyon/spin/sdk/go/http"
)

func main() {
	spin.HandleRequest(handle)
}

func handle(w http.ResponseWriter, r *http.Request) {
	filename := os.Getenv("FAVICON_PATH")
	if filename == "" {
		filename = "/favicon.ico"
	}
	mediaType := "image/vnd.microsoft.icon"

	// Open the file at /favicon and print it.
	var input, err = os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %s: %s\n", filename, err)
		send404(w)
		return
	}

	w.Header().Add("content-type", mediaType)
	if _, e := io.Copy(w, input); e != nil {
		fmt.Fprintf(os.Stderr, "Error writing file %s: %s\n", filename, e)
	}
}

func send404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Add("content-type", "text/plain")
	w.Write([]byte("Not Found"))
}
