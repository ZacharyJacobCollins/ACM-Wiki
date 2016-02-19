package main

import (
  "net/http"
  "regexp"
)

//Link to templates for wiki.  Remember that the templates are run from the executables path, NOT the path of the go file.
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

type Wiki struct {
    pages []Page
}

func NewWiki() Wiki {
	var wiki = Wiki {
		pages: make([]Page, 0),
	}
	return wiki
}

func (Wiki) Run() {
  home:=&Page{Title:"Home", Body: []byte("Welcome to the Acm Wiki")}
  home.save()
}
