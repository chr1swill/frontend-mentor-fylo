package main

import (
  "github.com/a-h/templ"
	"net/http"
	"log"
)

func main() {
  mux := http.NewServeMux();
	mux.Handle("/", templ.Handler(LandingPage("Baby")));

  log.Printf("serve starting on port :8080...");
	if err := http.ListenAndServe(":8080", mux);err != nil {
			log.Fatalf("failed to listen and serve on :8080 : %v", err);
	}
}
