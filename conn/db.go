package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/khanadnanxyz/konta/model"
	_ "github.com/lib/pq"
)

var (
	// Models describe models list for migration
	Models []interface{}
)

func main() {
	db, err := gorm.Open("postgres", "host=192.168.1.153 port=5432 user=postgres dbname=konta_db password=secret sslmode=disable")
	if err != nil {
		println(err.Error())
		println("error")
		return
	}
	println("success")
	defer db.Close()

	// add address model for migration
	Models = append(Models, &model.Role{})
	Models = append(Models, &model.User{})
	Models = append(Models, &model.Category{})
	Models = append(Models, &model.Question{})
	Models = append(Models, &model.Option{})
	Models = append(Models, &model.Answer{})

	db.AutoMigrate(Models...)
}

func init() {
	main()
}
