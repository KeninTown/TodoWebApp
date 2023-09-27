package usecase

import (
	"fmt"
	"strconv"
	"todos/internal/database/models"
	"todos/internal/dtos"
	"todos/internal/entites"
)

type Repository interface {
	CreateTodo(todo models.Todo) (*models.Todo, error)
	GetTodo(id uint) (*models.Todo, error)
	GetTodos() []models.Todo
	DeleteTodo(id uint)
	CompleteTodo(id uint, completed bool) (*models.Todo, error)
	UpdateTodo(updatedTodo *models.Todo) (*models.Todo, error)
}

type TodoUsecase struct {
	repo Repository
}

func New(repo Repository) TodoUsecase {
	return TodoUsecase{repo: repo}
}

func (tu TodoUsecase) Create(todo entites.Todo) (*models.Todo, error) {
	dbTodo := dtos.ToDbModel(todo)
	createdDbTodo, err := tu.repo.CreateTodo(dbTodo)
	if err != nil {
		return nil, err
	}

	return createdDbTodo, nil
}

func (tu TodoUsecase) Get(id string) (*models.Todo, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id")
	}

	todo, err := tu.repo.GetTodo(uint(intId))
	if err != nil {
		return nil, fmt.Errorf("todo is not exist")
	}

	return todo, nil
}

func (tu TodoUsecase) GetAll() []models.Todo {
	todos := tu.repo.GetTodos()
	return todos
}

func (tu TodoUsecase) Delete(id uint) {
	tu.repo.DeleteTodo(id)
}

func (tu TodoUsecase) Complete(id uint, completed bool) (*models.Todo, error) {
	todo, err := tu.repo.CompleteTodo(id, completed)
	if err != nil {
		return nil, fmt.Errorf("failed to complete todo: %s", err.Error())
	}
	return todo, nil
}

func (tu TodoUsecase) Update(todo entites.Todo) (*entites.Todo, error) {
	todoModel := dtos.ToDbModel(todo)

	updatedTodo, err := tu.repo.UpdateTodo(&todoModel)

	if err != nil {
		return nil, fmt.Errorf("failed to update todo")
	}

	updatedEntiteTodo := dtos.ToEntite(updatedTodo)
	return &updatedEntiteTodo, nil
}
