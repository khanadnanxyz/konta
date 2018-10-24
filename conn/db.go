package conn

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=http://192.168.1.153 port=5432 user=postgres dbname=konta_db password=secret")
	if err != nil {
		println(err.Error())
		println("errir")
		return
	}
	println("success")
	defer db.Close()
}

func init() {
	main()
}