package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"todos/internal/database/models"
	"todos/internal/dtos"
	"todos/internal/entites"
	sl "todos/pkg/logger"

	"golang.org/x/exp/slog"
)

type TodoUsecase interface {
	Create(todo entites.Todo) (*models.Todo, error) //done
	Get(id string) (*models.Todo, error)            //done
	GetAll() []models.Todo                          //done
	Delete(id uint)                                 //done
	Complete(id uint, completed bool) (*models.Todo, error)
	Update(todo entites.Todo) (*entites.Todo, error)
}

func SendMainPage(log *slog.Logger, tu TodoUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("./static/main.html")
		if err != nil {
			log.Error("failed to read main page html file", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		todos := tu.GetAll()

		responseTodos := dtos.EntitesTodo(&todos)
		positions := []uint8{}
		for i := range responseTodos {
			if i%2 == 0 {
				positions = append(positions, 0)
			} else {
				positions = append(positions, 1)
			}
		}

		data := struct {
			Todos     []entites.Todo
			Positions []uint8
		}{
			Todos:     responseTodos,
			Positions: positions,
		}

		//Change data there
		if err := temp.Execute(w, &data); err != nil {
			log.Error("failed to send main page html file", sl.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func CreateTodo(log *slog.Logger, tu TodoUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todo entites.Todo

		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			log.Error("failed to decode request body", sl.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "failed to decode request body")
			return
		}

		createdTodo, err := tu.Create(todo)

		if err != nil {
			log.Error("failed to create todo", sl.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "failed to decode request body")
			return
		}

		createdTodoResponse := dtos.ToEntite(createdTodo)

		log.Info("todo created", slog.Any("id", createdTodo.Id), slog.String("title", createdTodo.Title))
		json.NewEncoder(w).Encode(&createdTodoResponse)
	}
}

func GetTodos(log *slog.Logger, tu TodoUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos := tu.GetAll()
		responseTodos := dtos.EntitesTodo(&todos)
		json.NewEncoder(w).Encode(&responseTodos)
	}
}

func GetTodo(log *slog.Logger, tu TodoUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		todo, err := tu.Get(id)

		if err != nil {
			log.Error("failed to get todo", sl.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "failed to get todo")
			return
		}

		responseTodo := dtos.ToEntite(todo)
		json.NewEncoder(w).Encode(&responseTodo)
	}
}

func DeleteTodo(log *slog.Logger, tu TodoUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type DeletedTodo struct {
			Id uint `json:"id"`
		}

		var deletedTodo DeletedTodo

		if err := json.NewDecoder(r.Body).Decode(&deletedTodo); err != nil {
			log.Error("faile to decode todo's id", sl.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "invalid id")
			return
		}

		tu.Delete(deletedTodo.Id)

		w.WriteHeader(http.StatusOK)
	}
}

func CompleteTodo(log *slog.Logger, tu TodoUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type CompletedTodo struct {
			Id        uint `json:"id"`
			Completed bool `json:"completed"`
		}

		var completedTodo CompletedTodo

		if err := json.NewDecoder(r.Body).Decode(&completedTodo); err != nil {
			log.Error("failed to decode todo's id", sl.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "invalid id")
			return
		}

		todo, err := tu.Complete(completedTodo.Id, completedTodo.Completed)

		if err != nil {
			log.Error("failed to complete todo", sl.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err.Error())
			return
		}

		responseTodo := dtos.ToEntite(todo)

		json.NewEncoder(w).Encode(&responseTodo)
	}
}

// TODO: update todo
func UpdateTodo(log *slog.Logger, tu TodoUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todo entites.Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			log.Error("failed to decode request body", sl.Error(err))
			fmt.Fprintf(w, "failed to decode request body")
			return
		}
		updatedTodo, err := tu.Update(todo)

		if err != nil {
			log.Error(err.Error())
			fmt.Fprintf(w, "failed to update todo")
			return
		}

		json.NewEncoder(w).Encode(updatedTodo)
	}
}
