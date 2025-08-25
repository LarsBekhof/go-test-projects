package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func onFatalErr(err any) {
	fmt.Println(err)
	os.Exit(1)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	index, err := os.ReadFile("index.html")

	if err != nil {
		onFatalErr(err)
	}

	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, string(index))
}

func serveAssets(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	path := r.RequestURI[1:]

	asset, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "text/css")
	fmt.Fprintf(w, string(asset))
}

func main() {
	http.HandleFunc("GET /", serveIndex)
	http.HandleFunc("GET /assets/{asset}", serveAssets)

	port := ":9000"

	fmt.Println("Server is running on port " + port)

	log.Fatal(http.ListenAndServe(port, nil))
}
