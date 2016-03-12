package main

import (
  "net/http"

  //Third party packages
  "github.com/ZacharyJacobCollins/Wiki/wiki"
  "github.com/ZacharyJacobCollins/Wiki/chat"
)

// var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
// //Make handler takes a typical handler func as an anon func, and ensures that only registered
// //paths/handlers, found in validPath variable can be used.

// func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		//Valid path refers to global up to, to only view/edit/save
// 		m := validPath.FindStringSubmatch(r.URL.Path)
// 		if m == nil {
// 			http.NotFound(w, r)
// 			return
// 		}
// 		fn(w, r, m[2])
// 	}
// }
/*
        Wiki handlers, probably should put wiki in front
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
 */

func main() {
  c := chat.NewChat();  c.Run(3);
  w := wiki.NewWiki();  w.Run();
  http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
  http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./html/welcome"))))
	http.ListenAndServe(":1337", nil)
}
