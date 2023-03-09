/*
Create time at 2023年3月9日0009下午 16:22:37
Create User at Administrator
*/

package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "root:gg123456@tcp(127.0.0.1:3306)/booksystem?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = d
}
func GetDb() *gorm.DB {
	return db
}
