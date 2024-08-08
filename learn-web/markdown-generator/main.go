package main

import (
	"net/http"

	"github.com/russross/blackfriday"
)

func main() {
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("./markdown-generator/www/public")))
	http.ListenAndServe(":8080", nil)
}

func GenerateMarkdown(w http.ResponseWriter, req *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(req.FormValue("body")))
	w.Write(markdown)
}
