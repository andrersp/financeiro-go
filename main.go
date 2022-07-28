package main

import (
	"fmt"
	"log"

	"github.com/andrersp/financeiro-go/src/config"
	"github.com/andrersp/financeiro-go/src/http"
	"github.com/andrersp/financeiro-go/src/infra/persistence"
	"github.com/joho/godotenv"
)

func init() {

	// Generate Hash
	// key := make([]byte, 64)

	// if _, erro := rand.Read(key); erro != nil {
	// 	log.Fatal(erro)
	// }
	// strB64 := base64.StdEncoding.EncodeToString(key)
	// fmt.Println(strB64)
	// Load Env
	if err := godotenv.Load(); err != nil {
		log.Println("No Env loades")
	}

}

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
