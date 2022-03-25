package library

import "gorm.io/gorm"

type Authors struct {
	gorm.Model
	Author string `gorm:"foreignKey:Author;references:Author;type:varchar(100);column:Author"`
}

type AuthorSlice []Authors
