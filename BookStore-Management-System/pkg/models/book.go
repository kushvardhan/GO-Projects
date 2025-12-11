package models

import (
	"github.com/jinzhu/gorm"
	"github.com/kushvardhan/GO-Projects/BookStore-Management-System/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.ConnectDB()
	db=config.GetDB()
	db.AutoMigrate(&Book{})
}