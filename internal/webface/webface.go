package webface

import (
	"embed"
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//go:embed main.css
var css string

//go:embed index.tmpl.html
var b embed.FS

func Start() {

	mux := http.NewServeMux()

	mux.HandleFunc("/main.css", cssHandler)
	mux.HandleFunc("/", HomeHandler)

	http.Handle("/", mux)
	fmt.Println("web server start on 127.0.0.1:8080")
	err := http.ListenAndServe("127.0.0.1:8080", mux)

	if err != nil {
		log.Fatalln(err)
	}
}

func cssHandler(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprint(writer, css)
	if err != nil {
		log.Fatalln(err)
	}
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {

	templ, err := template.ParseFS(b, "*.tmpl.html")
	if err != nil {
		log.Fatalln(err)
	}
	err = templ.Execute(writer, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
