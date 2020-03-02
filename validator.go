package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

type validateErr struct {
	fp  string
	err error
}

func (v validateErr) Error() string {
	return fmt.Sprintf("%s: %s", v.fp, v.err)
}

func (v validateErr) Unwrap() error {
	return v.err
}

type errorSet []error

func (e errorSet) Error() string {
	errsS := make([]string, len(e))
	for i, err := range e {
		errsS[i] = err.Error()
	}
	return strings.Join(errsS, "\n")
}

func validatePath(root string) error {
	var (
		errs = errorSet{}
		m    sync.Mutex
		wg   sync.WaitGroup
	)
	if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(path)
			return err
		}
		if !info.IsDir() {
			return nil
		}
		var buf bytes.Buffer
		if err := populateHTML(&buf, path); err != nil {
			return err
		}
		links, err := getLinks(&buf)
		if err != nil {
			return err
		}
		wg.Add(1)
		go func(links []string) {
			defer wg.Done()
			if err := validateLinks(links); err != nil {
				m.Lock()
				errs = append(errs, validateErr{err: err, fp: path})
				m.Unlock()
			}
		}(links)
		return nil
	}); err != nil {
		return err
	}
	wg.Wait()
	if len(errs) != 0 {
		return errs
	}
	return nil
}

func getLinks(src io.Reader) ([]string, error) {
	links := []string{}

	tokenizer := html.NewTokenizer(src)
	err := tokenizer.Err()
	for err == nil {
		_, hasAttr := tokenizer.TagName()
		if hasAttr {
			k, v, _ := tokenizer.TagAttr()
			if bytes.EqualFold(k, []byte("href")) {
				// href tag
				links = append(links, string(v))
			}
		}
		tokenizer.Next()
		err = tokenizer.Err()
	}
	if err != io.EOF {
		return links, err
	}
	return links, nil
}

func validateLinks(links []string) error {
	for _, link := range links {
		u, err := url.Parse(link)
		if err != nil {
			return fmt.Errorf("error parsing '%s': %w", link, err)
		}
		if strings.HasPrefix(u.Host, "localhost") || strings.HasPrefix(u.Host, "example") {
			continue
		}

		resp, err := http.DefaultClient.Head(link)
		if err != nil {
			return fmt.Errorf("error retrieving '%s': %w", link, err)
		}
		if resp.StatusCode >= 400 {
			return fmt.Errorf("http HEAD '%s' status code: %d", link, resp.StatusCode)
		}

		if err := resp.Body.Close(); err != nil {
			return fmt.Errorf("error closing resp body: %w", err)
		}
	}

	return nil
}
