package main

import (
  "context"
	"net/http"
	"fmt"
	"log"
	"os"
)

func main() {
	dirName := "html";
	_, err := os.Stat(dirName);
	if err != nil {
    err := os.Mkdir(dirName, os.FileMode(0755));
		if err != nil {
			log.Fatalf("failed to create directory '%s/': %v\n");
		}
	}
		
	f, err := os.OpenFile(fmt.Sprintf("%s/hello.html", dirName), os.O_RDWR|os.O_CREATE, 0644);
	if err != nil {
		log.Fatalf("failed to create file '%s/hello.html': %v\n", dirName, err);
	}
	defer f.Close();

  mux := http.NewServeMux();
	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
			err = LandingPage("John").Render(context.Background(), w);
			if err != nil {
			  log.Fatalf("failed to write output file: %v", err);
			}
	});

  log.Printf("serve starting on port :8080...");
	if err := http.ListenAndServe(":8080", mux);err != nil {
			log.Fatalf("failed to listen and serve on :8080 : %v", err);
	}
}
