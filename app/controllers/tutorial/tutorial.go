package tutorial

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/tomotakashimizu/todoapp-go/config"
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

var templates = template.Must(template.ParseFiles("app/views/tutorial/view.html", "app/views/tutorial/edit.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// t, err := template.ParseFiles("app/views/" + tmpl + ".html")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	// // w に p の値を使って html を表示
	// t.Execute(w, p)

	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		// http://localhost:8080/view/test
		// fmt.Println(m, r.URL.Path)
		// [/view/test view test] /view/test
		if m == nil {
			http.NotFound(w, r)
			fmt.Println("NotFound")
			return
		}
		fn(w, r, m[2])
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// URL の /view/ 以降の文字列を title に代入
	// title := r.URL.Path[len("/view/"):]

	// title.txt の file があれば p に情報を代入
	p, err := LoadPage(title)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	// w に 値を書き込んで表示
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, _ *http.Request, title string) {
	// title := r.URL.Path[len("/edit/"):]
	p, err := LoadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body") //この値が取れなくてもエラーは起きず空の値が入る
	p := &Page{Title: title, Body: []byte(body)}
	err := p.SavePage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+p.Title, http.StatusFound)
}

func StartTutorialServer() error {
	// URL /view/ の処理を関数で定義
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	// サーバを起動
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
