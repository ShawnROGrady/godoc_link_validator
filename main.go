package main

import (
	"flag"
	"log"
	"regexp"
)

func main() {
	checkExpr := flag.String("check", ".", "a regex indicating the URL host(s) which should be checked")
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatal("must provide root of path to validate")
	}
	re, err := regexp.Compile(*checkExpr)
	if err != nil {
		log.Fatalf("error compiling provided check expression: %s", err)
	}

	if err := validatePath(flag.Args()[0], re); err != nil {
		log.Fatal(err)
	}
}
