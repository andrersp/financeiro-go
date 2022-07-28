package persistence

import (
	"fmt"
	"log"

	"github.com/andrersp/financeiro-go/src/domain/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DBConf() (db *gorm.DB, err error) {
	fmt.Println("")
	db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)

	}

	err = db.AutoMigrate(&entity.User{})

	if err != nil {
		log.Fatal(err)
	}

	return
}
