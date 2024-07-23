package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/assci-art", AsciiHandler)
	http.HandleFunc("/h", HomeHandler)

	log.Println("http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("errur")
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Internalserver error", 500)
	}
	fmt.Println("banner, text")

	tmpl.Execute(w, nil)
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	banner := r.FormValue("banner")
	fmt.Println(banner)
}
