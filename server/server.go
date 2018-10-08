package server

import (
	"fmt"
	"net/http"
	"regexp"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Home Page.")
}

func ImagesHandler(w http.ResponseWriter, r *http.Request) {
	pathValidator := regexp.MustCompile("/images/([a-zA-Z0-9]{12})/((\\d+)x(\\d+))$")
	m := pathValidator.FindStringSubmatch(r.URL.Path)
	if len(m) == 0 {
		fmt.Fprintf(w, "can't compile regexp, no Id is found\nURL.Path: %v", r.URL.Path)
		return
	}
	fmt.Fprintf(w, "Found Id for image %v.\n", m[1])
}
