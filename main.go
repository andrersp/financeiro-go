package main

import (
	"fmt"
	"log"

	"github.com/andrersp/financeiro-go/src/config"
	"github.com/andrersp/financeiro-go/src/http"
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

	r := http.StartServer(*services)
	r.Run()

}
