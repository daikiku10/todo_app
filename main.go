package main

import (
	"fmt"
	"todo_app/app/models"
	// "log"
	// "todo_app/config"
)

func main() {
	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(1)

	// fmt.Println(u)

	// u.Name = "kudo"
	// u.Email = "kudo@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()

	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// user, _ := models.GetUser(2)
	// user.CreateTodo("First Todo")

	t, _ := models.GetTodo(1)
	fmt.Println(t)
}