package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB

func Connect() {
	db, err := gorm.Open("postgres", "host=192.168.1.153 port=5432 user=postgres dbname=konta_db password=secret sslmode=disable")
	if err != nil {
		println(err.Error())
		println("error")
		return
	}
	println("success")
	Db = db

}
