package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func CreateSQLiteConnection() (err error) {

	if dbConn != nil {
		CloseSQLiteConnection(dbConn)
	}

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		return
	}

	dbConn = db
	return
}

func ConnectSQLite() (db *gorm.DB, err error) {
	sqlDb, err := dbConn.DB()

	if err != nil {
		return
	}

	if err = sqlDb.Ping(); err != nil {
		return
	}
	db = dbConn
	return
}

func CloseSQLiteConnection(conn *gorm.DB) {
	db, err := conn.DB()
	if err != nil {
		return
	}

	defer db.Close()
}
