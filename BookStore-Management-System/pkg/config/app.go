package config

import (
	"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func ConnectDB() {
	d, err := gorm.Open("mysql", "root:@/BookStore?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
