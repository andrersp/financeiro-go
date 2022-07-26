package main

import (
	"fmt"
	"log"

	"github.com/andrersp/financeiro-go/src/config"
	"github.com/andrersp/financeiro-go/src/domain/entity"
	"github.com/andrersp/financeiro-go/src/infra/persistence"
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

	user := entity.User{
		FirstName: "Andre F",
		LastName:  "Luis",
		Password:  "andreluis",
	}

	r, err := services.User.SaveUser(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*r)

}
