package main

import (
  "log"
	"net/http"
	"text/template"
)

func main () {
	var err error;
	var mux *http.ServeMux;
	var tmpl *template.Template;

	tmpl, err = tmpl.ParseGlob("./html/");
	if err != nil {
		panic(err);
	}

	mux = http.NewServeMux();
	log.Println(mux, tmpl);
}
