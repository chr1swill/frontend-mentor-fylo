package main

import (
  "context"
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

	err = Hello("John").Render(context.Background(), f);
	if err != nil {
		log.Fatalf("failed to write output file: %v", err);
	}
}
