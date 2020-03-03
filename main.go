package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("must provide root of path to validate")
	}
	if err := validatePath(os.Args[1]); err != nil {
		log.Fatal(err)
	}
}
