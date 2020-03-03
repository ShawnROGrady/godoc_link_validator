package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
)

type reFlag struct {
	re *regexp.Regexp
}

func (r *reFlag) String() string {
	if r.re == nil {
		return ""
	}
	return r.re.String()
}

func (r *reFlag) Set(val string) error {
	re, err := regexp.Compile(val)
	if err != nil {
		return fmt.Errorf("error compiling regex: %s", err)
	}
	r.re = re
	return nil
}

func main() {
	var (
		checkRe  = reFlag{}
		ignoreRe = reFlag{}
	)
	flag.Var(&checkRe, "check", "A regex indicating the URL host(s) which should be checked. If empty all URLs will be checked.")
	flag.Var(&ignoreRe, "ignore", "A regex indicating the URL host(s) which should be ignored. If empty no URLs will be ignored.")
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatal("must provide root of path to validate")
	}

	v := validator{
		checkRe:  checkRe.re,
		ignoreRe: ignoreRe.re,
	}
	if err := v.validatePath(flag.Args()[0]); err != nil {
		log.Fatal(err)
	}
}
