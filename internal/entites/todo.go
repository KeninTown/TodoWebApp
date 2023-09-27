package entites

import "time"

type Todo struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	Type        []string  `json:"type"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	Deadline    time.Time `json:"deadline"`
}
