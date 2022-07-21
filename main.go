package main

import (
	"fmt"
	"log"

	"github.com/tomotakashimizu/todoapp-go/app/models"
)

func main() {
	fmt.Println(models.Db)

	t, err := models.GetAllTodos()
	if err != nil {
		log.Fatalf("failed to GetAllTodos: %s", err.Error())
	}
	fmt.Println(t)
}
