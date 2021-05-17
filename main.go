package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

//go:embed templates
var content embed.FS

//go:embed static
var staticContent embed.FS

func init() {
	fmt.Println(content.ReadDir("."))
	tpl = template.Must(template.ParseFS(content, "templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/second", second)
	http.Handle("/static/", http.FileServer(http.FS(staticContent)))
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func second(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "second.gohtml", nil)
}
