package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/khanadnanxyz/konta/config"
)

var Db *gorm.DB

func Connect() {
	host := config.AppConfig.GetString("postgres.host")
	port := config.AppConfig.GetString("postgres.port")
	username := config.AppConfig.GetString("postgres.username")
	password := config.AppConfig.GetString("postgres.password")
	database := config.AppConfig.GetString("postgres.db")
	db, err := gorm.Open("postgres", "host=" + host + " port=" + port +
		" user="+ username +" dbname=" + database + " password=" + password + " sslmode=disable")
	if err != nil {
		println(err.Error())
		println("errors")
		return
	}
	println("success")
	Db = db
	//db.AutoMigrate(&model.Question2{})
	//db.AutoMigrate(&model.Role{})
	//db.AutoMigrate(&model.User{})
	//db.AutoMigrate(&model.Category{})
	//db.AutoMigrate(&model.Question{})
	//db.AutoMigrate(&model.Option{})
	//db.AutoMigrate(&model.Answer{})

}
