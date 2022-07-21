package main

import (
	"log"

	"github.com/tomotakashimizu/todoapp-go/app/controllers"
)

func main() {
	log.Fatalln(controllers.StartMainServer())
}
