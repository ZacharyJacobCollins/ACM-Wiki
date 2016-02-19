package main

import (
  "net/http"

  //Third party imports
	"github.com/ZacharyJacobCollins/Wiki/chat"
  "github.com/ZacharyJacobCollins/Wiki/wiki"
)

//Make handler takes a typical handler func as an anon func, and ensures that only registered
//paths/handlers, found in validPath variable can be used.
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Valid path refers to global up to, to only view/edit/save
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func RunServer() {
  c := chat.NewChat();  c.Run();
  w := wiki.NewWiki();  w.Run();
  http.HandleFunc("/view/", makeHandler(viewHandler))
  http.HandleFunc("/edit/", makeHandler(editHandler))
  http.HandleFunc("/save/", makeHandler(saveHandler))
  
	http.ListenAndServe(":1337", nil)
}
