package wiki

import (
  "net/http"
)

type Wiki struct {
    Pages []Page
}

func (Wiki)Run() {
  home:=&Page{Title:"Home", Body: []byte("Welcome to the Acm Wiki")}
  home.save()
  http.HandleFunc("/view/", makeHandler(viewHandler))
  http.HandleFunc("/edit/", makeHandler(editHandler))
  http.HandleFunc("/save/", makeHandler(saveHandler))
}
