package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)

func init() {
	spinhttp.Handle(favicon)
}

func favicon(w http.ResponseWriter, r *http.Request) {
	filename := os.Getenv("FAVICON_PATH")
	if filename == "" {
		filename = "/favicon.ico"
	}
	mediaType := "image/vnd.microsoft.icon"

	var input, err = os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %s: %s\n", filename, err)

		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Not found")
		return
	}

	w.Header().Set("content-type", mediaType)
	if _, e := io.Copy(w, input); e != nil {
		fmt.Fprintf(os.Stderr, "Error writing file %s: %s\n", filename, e)
	}
}

func main() {}
