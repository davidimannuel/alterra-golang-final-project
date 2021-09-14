package label

import (
	"gorm.io/gorm"
)

var tableName string = "labels"

type LabelModel struct {
	gorm.Model
	UserID uint
	Name   string
}

func (model *LabelModel) TableName() string {
	return tableName
}
