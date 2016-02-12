package main

import (
	"github.com/ZacharyJacobCollins/Wiki/chat"
  "github.com/ZacharyJacobCollins/Wiki/wiki"
	"net/http"
)

func main() {
  http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	c := chat.NewChat()
	c.Run()
  w := wiki.NewWiki()
  w.Run()
	http.ListenAndServe(":1337", nil)
}
