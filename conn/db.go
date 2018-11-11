package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var Db *gorm.DB

func Connect() {
	host := viper.GetString("postgres.host")
	port := viper.GetString("postgres.port")
	username := viper.GetString("postgres.username")
	password := viper.GetString("postgres.password")
	database := viper.GetString("postgres.db")
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
