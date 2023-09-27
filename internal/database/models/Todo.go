package models

import (
	"time"
)

type Todo struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	Created_At  time.Time
	Deadline    time.Time
}
