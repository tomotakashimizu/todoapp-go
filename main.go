package main

import (
	"fmt"

	"github.com/tomotakashimizu/todoapp-go/app/models"
)

func main() {
	fmt.Println(models.Db)

	u := &models.User{Name: "testuser4", Email: "test4@example.com", Password: "test4password"}
	fmt.Println(u)

	u.CreateUser()
}
