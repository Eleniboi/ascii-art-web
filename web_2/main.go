package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/ascii", asciiHandler)
	mux.HandleFunc("/switch", switchHandler)

	fmt.Println("server is live on port 8080....")

	http.ListenAndServe(":8080", mux)

}
