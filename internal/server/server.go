package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"todos/internal/server/handlers"
	"todos/internal/server/middleware"
	sl "todos/pkg/logger"

	"github.com/gofiber/fiber/v2/log"
	"github.com/gorilla/mux"
	"golang.org/x/exp/slog"
)

type Server struct {
	Usecase handlers.TodoUsecase
	Log     *slog.Logger
	Port    string
}

func New(port string, usecase handlers.TodoUsecase, log *slog.Logger) *Server {
	return &Server{
		Usecase: usecase,
		Log:     log,
		Port:    port,
	}
}

func (s Server) Run(ctx context.Context) {
	router := mux.NewRouter()

	//handle static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	// router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	//TODO: send html, css, js files
	router.HandleFunc("/", middleware.Logger(s.Log, handlers.SendMainPage(s.Log, s.Usecase))).Methods("GET")

	//TODO: get todo by query id
	router.HandleFunc("/todo", middleware.Logger(s.Log, handlers.GetTodo(s.Log, s.Usecase))).Methods("GET")

	//TODO: create todo
	router.HandleFunc("/todos", middleware.Logger(s.Log, handlers.CreateTodo(s.Log, s.Usecase))).Methods("POST")

	//TODO: get todos
	router.HandleFunc("/todos", middleware.Logger(s.Log, handlers.GetTodos(s.Log, s.Usecase))).Methods("GET")

	//TODO: delete todo
	router.HandleFunc("/todos", middleware.Logger(s.Log, handlers.DeleteTodo(s.Log, s.Usecase))).Methods("DELETE")

	//TODO: complete todo
	router.HandleFunc("/complete", middleware.Logger(s.Log, handlers.CompleteTodo(s.Log, s.Usecase))).Methods("PUT")

	//update todo
	router.HandleFunc("/todos", middleware.Logger(s.Log, handlers.UpdateTodo(s.Log, s.Usecase))).Methods("PUT")
	srv := http.Server{
		Addr:    s.Port,
		Handler: router,
	}

	go func() {
		log.Info(fmt.Sprintf("server listen on port %s", s.Port))
		if err := srv.ListenAndServe(); err != nil {
			log.Error("failed to listen server", sl.Error(err))
		}
	}()

	//gracefull shutdown
	<-ctx.Done()
	log.Info("starting to shutdown server...")
	shutDownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutDownCtx); err != nil {
		log.Error("failed to shutdown server gracefully", sl.Error(err))
	}
	log.Info("server successfully shutdown")
}
