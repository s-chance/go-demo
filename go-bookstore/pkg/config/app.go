package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "entropy:entropy@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local&timeout=10s")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
