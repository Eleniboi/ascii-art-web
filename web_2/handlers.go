package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type pageData struct {
	Text   string
	Result string
	Banner string
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 Method Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl.Execute(w, nil)
}


func asciiHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/ascii" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {

		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	result, err := Artbuilder(text, banner)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	Data := pageData{

		Text:   text,
		Result: result,
		Banner: "",
	}

	tmpl.Execute(w, Data)
}

func switchHandler(w http.ResponseWriter, r *http.Request) {

	text := r.FormValue("text")
	banner := r.FormValue("banner")
	fmt.Println(text)
	fmt.Println(banner)

	result, err := Artbuilder(text, banner)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	Data := pageData{

		Text:   text,
		Result: result,
	}

	tmpl.Execute(w, Data)
}
