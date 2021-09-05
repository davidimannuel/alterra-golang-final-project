package repository

import "gorm.io/gorm"

type noteRepository struct {
	conn *gorm.DB
}
