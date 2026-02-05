package main

import (
  "github.com/a-h/templ"
	"net/http"
	"log"
)

func main() {
  mux := http.NewServeMux();
	fs := http.FileServer(http.Dir("./static"));

	mux.Handle("/static/", http.StripPrefix("/static/", fs));
	mux.Handle("/", templ.Handler(HomePage()));

  log.Printf("serve starting on port :8080...");
	if err := http.ListenAndServe(":8080", mux);err != nil {
			log.Fatalf("failed to listen and serve on :8080 : %v", err);
	}
}
