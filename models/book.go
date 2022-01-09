package models

import (
	"time"
)

type Book struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Title     string    `json:"title" binding:"required,min=2,max=100" gorm:"type:varchar(100)"`
	Author    Person    `json:"author" binding:"required" gorm:"foreignkey:AuthorID"`
	AuthorID  uint64    `json:"-"`
	WrittenIn string    `json:"written_in" gorm:"type:varchar(5)"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
