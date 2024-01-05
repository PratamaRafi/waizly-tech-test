package models

import (
	"time"
)

type (
	// customer
	Task struct {
		ID          uint      `gorm:"primary_key" json:"task_id"`
		CustomerID  uint      `gorm:"foreign_key" json:"customer_id"`
		Description string    `gorm:"not null" json:"description"`
		DueDate     time.Time `gorm:"default:null" json:"due_date"`
		IsDone      bool      `gorm:"default:false" json:"is_done type:boolean"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
