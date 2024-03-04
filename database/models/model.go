package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id    uint64 `gorm:"primary_key;auto_increment;column:id"`
	Name  string `gorm:"type:varchar(255);column: name"`
	Email string `gorm:"uniqueIndex;type:varchar(255);column: email1"`
}
