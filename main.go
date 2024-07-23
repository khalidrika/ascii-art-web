package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/assci-art", AsciiHandler)
	http.HandleFunc("/", HomeHandler)
	css := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css", css))
	log.Println("http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("errur")
	}
}

type Data struct {
	Font  string
	Fonts map[string]string
	Out   string
	Text  string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// out := Fonts{}
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Internalserver error", http.StatusInternalServerError)
	}
	// fmt.Println("banner, text")
	// out.Fonts = append(out.Fonts, "stander")
	// out.Fonts = append(out.Fonts, "shadow")

	// Example data
	data := Data{
		Font: "Ariag gvytgtygbyubyuhbuybuhl", // Selected font
		Fonts: map[string]string{
			"standard":   "standard",
			"thinkertoy": "thinkertoy",
			"shadow":     "shadow",
		},
		Out:  "test",
		Text: "",
	}
	err = (tmpl.Execute(w, data))
	if err != nil {
		fmt.Println(err)
	}
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	banner := r.FormValue("font")
	text := r.FormValue("text")
	out := text
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
	}
	data := Data{
		Font: "Aritgfvtyfvytftygfytgytgytgytivfyti", // Selected font
		Fonts: map[string]string{
			"standard":   "standard",
			"thinkertoy": "thinkertoy",
			"shadow":     "shadow",
			"arob":    "arob",
		},
		Out:  out,
		Text: text,
	}
	tmpl.Execute(w, data)
	fmt.Println(banner)
}
