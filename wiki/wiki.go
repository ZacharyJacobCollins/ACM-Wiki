package wiki

import (
	"net/http"
	"regexp"
)

//Link to templates for wiki.  Remember that the templates are run from the executables path, NOT the path of the go file.
var validPath = regexp.MustCompile("^/wiki/(edit|save|view)/([a-zA-Z0-9]+)$")

type Wiki struct {
	pages []Page
}

func NewWiki() Wiki {
	var wiki = Wiki{
		pages: make([]Page, 0),
	}
	return wiki
}

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

func (Wiki) Run() {
	home := &Page{Title: "Home", Body: []byte("Welcome to the Acm Wiki")}
	home.save()
	http.HandleFunc("/wiki/view/", makeHandler(viewHandler))
	http.HandleFunc("/wiki/edit/", makeHandler(editHandler))
	http.HandleFunc("/wiki/save/", makeHandler(saveHandler))
}
