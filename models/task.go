package models

import "time"

type Task struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Name      string    `json:"name" binding:"required" gorm:"type:varchar(120)"`
	Status    string    `json:"status" gorm:"type:varchar(50);default:'closed'"`
	Priority  uint64    `json:"priority" gorm:"default:1"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	CreatedBy Person    `json:"created_by" binding:"required" gorm:"foreignkey:CreatorID"`
	CreatorID uint64    `json:"-"`
	DueDate   string    `json:"dueDate"`
}
