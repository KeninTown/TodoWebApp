package dtos

import (
	"fmt"
	"todos/internal/database/models"
	"todos/internal/entites"
)

func ToDbModel(todo entites.Todo) models.Todo {
	var todoType string

	for i, t := range todo.Type {
		if i != len(todo.Type)-1 {
			todoType += t + " "
		} else {
			todoType += t
		}
	}
	fmt.Printf("%+v", todo.Type)
	return models.Todo{
		Id:          todo.Id,
		Title:       todo.Title,
		Type:        todoType,
		Description: todo.Description,
		Deadline:    todo.Deadline,
	}
}
