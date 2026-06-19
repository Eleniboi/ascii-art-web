package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/ascii", asciiHandler)

	fmt.Println("server is live on port 8080....")
	http.ListenAndServe(":8080", mux)

}
