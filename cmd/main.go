package main

import (
  "context"
	"log"
	"os"
)

func main() {
	f, err := os.Create("hello.html");
	if err != nil {
		log.Fatalf("failed to create output file: %v", err);
	}

	err = Hello("John").Render(context.Background(), f);
	if err != nil {
		log.Fatalf("failed to write output file: %v", err);
	}
}
