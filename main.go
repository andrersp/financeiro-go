package main

import (
	"fmt"
	"log"

	"github.com/andrersp/financeiro-go/src/config"
	"github.com/andrersp/financeiro-go/src/infra/persistence"
	"github.com/andrersp/financeiro-go/src/interfaces"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Start")

	err := config.CreateSQLiteConnection()

	if err != nil {
		log.Fatal(err)
	}

	db, err := config.ConnectSQLite()

	if err != nil {
		log.Fatal(err)
	}

	services, err := persistence.NewRepositories(db)

	services.AutoMigrate()

	users := interfaces.NewUserHandler(services.User)

	r := gin.Default()
	r.POST("/user", users.SaveUser)

	r.Run()

	// user := entity.User{
	// 	FirstName: "Andre F",
	// 	LastName:  "Luis",
	// 	Password:  "andreluis",
	// }

	// r, err := services.User.SaveUser(user)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(*r)

}
