package wiki

import (
  "net/http"
)

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
  http.HandleFunc("/view/", makeHandler(viewHandler))
  http.HandleFunc("/edit/", makeHandler(editHandler))
  http.HandleFunc("/save/", makeHandler(saveHandler))
}
