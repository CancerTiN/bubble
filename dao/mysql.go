package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	db, err := InitMySQL()
	if err != nil {
		panic(err)
	}
	DB = db
}

func InitMySQL() (db *gorm.DB, err error) {
	dialect := "mysql"
	arg := "root:root@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(dialect, arg)
	if err != nil {
		panic(err)
	}
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	return
}

func CloseMySQL() {
	_ = DB.Close()
}
