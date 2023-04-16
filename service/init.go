package service

import (
	"github.com/shplume/ygo-cards/model"
	"gorm.io/gorm"
)

var (
	cards  = map[string]interface{}{}
	dbConn *gorm.DB
)

func init() {
	dbConn = model.GetDBConnection()
	load()
}
