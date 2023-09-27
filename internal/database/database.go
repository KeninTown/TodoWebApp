package database

import (
	"errors"
	"fmt"
	"todos/internal/config"
	"todos/internal/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBInstance struct {
	db *gorm.DB
}

func New(cfgDb config.Db) (*DBInstance, error) {
	op := "database.New()"
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable TimeZone=Europe/Moscow",
		cfgDb.User, cfgDb.Password, cfgDb.DbName, cfgDb.Host, cfgDb.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("%s: failed connect to db: %s", op, err.Error())
	}
	return &DBInstance{
		db: db,
	}, nil
}

// TODO: create todo
func (dbi DBInstance) CreateTodo(todo models.Todo) (*models.Todo, error) {
	op := "DBInstance.CreateTodo()"
	dbi.db.Create(&todo)
	if todo.Id == 0 {
		return nil, fmt.Errorf("%s: failed to create todo: %+v", op, todo)
	}

	return &todo, nil
}

// TODO: get todo
func (dbi DBInstance) GetTodo(id uint) (*models.Todo, error) {
	var todo models.Todo

	if err := dbi.findTodo(id, &todo); err != nil {
		return nil, err
	}
	return &todo, nil
}

// TODO: get todos
func (dbi DBInstance) GetTodos() []models.Todo {
	var todo []models.Todo
	dbi.db.Find(&todo)
	return todo
}

// TODO: delete todo
func (dbi DBInstance) DeleteTodo(id uint) {
	var todo models.Todo
	dbi.db.Delete(&todo, "id = ?", id)
}

// TODO: complete todo
func (dbi DBInstance) CompleteTodo(id uint, completed bool) (*models.Todo, error) {
	var todo models.Todo

	if err := dbi.findTodo(id, &todo); err != nil {
		return nil, err
	}

	todo.Completed = completed
	dbi.db.Save(&todo)

	return &todo, nil
}

// TODO: update todo
func (dbi DBInstance) UpdateTodo(updatedTodo *models.Todo) (*models.Todo, error) {
	var todo models.Todo

	if err := dbi.findTodo(updatedTodo.Id, &todo); err != nil {
		return nil, err
	}

	todo.Title = updatedTodo.Title
	todo.Description = updatedTodo.Description
	todo.Type = updatedTodo.Type
	todo.Completed = updatedTodo.Completed

	dbi.db.Save(&todo)

	return &todo, nil
}

func (dbi DBInstance) findTodo(id uint, todo *models.Todo) error {
	dbi.db.Find(todo, "id = ?", id)
	if todo.Id == 0 {
		return errors.New("todo is not exist")
	}

	return nil
}
