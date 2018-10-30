package repository

import (
	"github.com/khanadnanxyz/konta/conn"
	"github.com/khanadnanxyz/konta/model"
)

func CreateQuestion2(question *model.Question2)  {
	f := db.Db.NewRecord(question) // => returns `true` as primary key is blank
	db.Db.Create(&question)
	println(f)
}
