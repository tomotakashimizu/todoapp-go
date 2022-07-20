package main

import (
	"fmt"

	"github.com/tomotakashimizu/todoapp-go/app/models"
)

func main() {
	fmt.Println(models.Db)

	u, err := models.GetUser(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}
