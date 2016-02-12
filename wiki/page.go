package wiki

import (
	"errors"
	"html/template"
	"io/ioutil"
	"net/http"
)

//Link to templates for wiki.  Remember that the templates are run from the executables path, NOT the path of the go file.
var templates = template.Must(template.ParseFiles("./wiki/templates/view.html", "./wiki/templates/edit.html"))

type Page struct {
	Title string
	Body  []byte
}

//Saves a page
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//Loads the saved page
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}, nil
}

//Renders the current template
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Gets the title for the wiki post
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}
