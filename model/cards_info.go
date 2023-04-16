package model

import (
	"gorm.io/datatypes"
)

type CardInfo struct {
	ID      uint `gorm:"primarykey"`
	Details datatypes.JSON
}

func (c *CardInfo) TableName() string {
	return database + ".card_infos"
}
