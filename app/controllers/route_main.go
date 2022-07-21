package controllers

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) SavePage() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles("app/views/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// w に p の値を使って html を表示
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// URL の /view/ 以降の文字列を title に代入
	title := r.URL.Path[len("/view/"):]
	// title.txt の file があれば p に情報を代入
	p, err := LoadPage(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// w に 値を書き込んで表示
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := LoadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}
