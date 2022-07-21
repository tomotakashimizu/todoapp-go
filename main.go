package main

import (
	"fmt"
	"log"

	"github.com/tomotakashimizu/todoapp-go/app/models"
)

func main() {
	fmt.Println(models.Db)

	t := models.Todo{Content: "Todofor2_2", UserID: 2}
	fmt.Println(t)

	err := t.CreateTodo()
	if err != nil {
		log.Fatalf("failed to Create Todo: %s", err.Error())
	}
}
