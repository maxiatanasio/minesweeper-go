package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

func ConfigDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:12345678@tcp(localhost)/minescrapper")
	if err != nil {
		print("There was a problem connecting to the DB", err.Error())
		os.Exit(1)
	}

	defer db.Close()

	return db
}
