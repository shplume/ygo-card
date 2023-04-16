package model

import (
	"fmt"

	"github.com/shplume/ygo-cards/util/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	conn     *gorm.DB
	database = config.Getstring("database")
)

func GetDBConnection() *gorm.DB {
	return conn
}

func init() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=%s&parseTime=True&loc=Local",
		config.Getstring("userName"), config.Getstring("password"),
		config.Getstring("host"), config.Getstring("port"),
		config.Getstring("charset"))

	conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	conn.Exec("CREATE DATABASE IF NOT EXISTS " + database + " DEFAULT CHARACTER SET " + config.Getstring("charset"))

	conn.AutoMigrate(&CardInfo{})
}
