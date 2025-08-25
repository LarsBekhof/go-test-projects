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

func main() {
	http.Handle("/", http.FileServer(http.Dir("assets")))

	port := ":9000"

	fmt.Println("Server is running on port " + port)

	log.Fatal(http.ListenAndServe(port, nil))
}
