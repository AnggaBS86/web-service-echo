package models

import (
	"time"
)

type Task struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	AssingedTo string    `json:"assigned_to"`
	Task       string    `json:"task"`
	Deadline   time.Time `json:"deadline"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
