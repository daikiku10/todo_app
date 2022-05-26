package main

import (
	"fmt"
	"todo_app/app/models"
	// "log"
	// "todo_app/config"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Panicln("test")

	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	u, _ := models.GetUser(1)

	fmt.Println(u)

	u.Name = "kudo"
	u.Email = "kudo@example.com"
	u.UpdateUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)
}