package main

import (
	"fmt"

	"github.com/tomotakashimizu/todoapp-go/app/controllers"
)

func main() {
	p1 := controllers.Page{Title: "test", Body: []byte("This is a sample page.")}
	p1.SavePage()

	p2, _ := controllers.LoadPage(p1.Title)
	fmt.Println(string(p2.Body))
}
