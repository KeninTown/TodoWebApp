package dtos

import (
	"strings"
	"todos/internal/database/models"
	"todos/internal/entites"
)

func ToEntite(todo *models.Todo) entites.Todo {
	todoType := strings.Split(todo.Type, " ")

	return entites.Todo{
		Id:          todo.Id,
		Title:       todo.Title,
		Type:        todoType,
		Description: todo.Description,
		Completed:   todo.Completed,
	}
}

func EntitesTodo(todos *[]models.Todo) []entites.Todo {
	var entitysTodos []entites.Todo

	for _, todo := range *todos {
		entityTodo := ToEntite(&todo)
		entitysTodos = append(entitysTodos, entityTodo)
	}

	return entitysTodos
}
