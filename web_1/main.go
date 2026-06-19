package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type file struct {
	Title  string
	Result string
	Banner     string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("template/index.html")

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	
	banner := r.FormValue("banner")

	Data := file{

		Title:  "Ascii-Art-Web",
		Result: "",
		Banner: banner,
	}
	tmpl.Execute(w, Data)
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/ascii-art" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	result, err := ArtBuilder(text, banner)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("template/index.html")

	if err != nil {
		log.Fatal("error parsing file ")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	Data := file{

		Title:  "Ascii-Art-Web",
		Result: result,
	}

	tmpl.Execute(w, Data)

}

func main() {

	fs := http.FileServer(http.Dir("./static/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiHandler)

	fmt.Println("server is now live on port 3000...")

	http.ListenAndServe(":3000", nil)
}
