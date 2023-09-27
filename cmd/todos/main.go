package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"todos/internal/config"
	"todos/internal/database"
	"todos/internal/logger"
	"todos/internal/server"
	"todos/internal/usecase"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	defer cancel()

	//replace to env in docker
	cfgPath := os.Getenv("CONFIG_PATH")

	//read config file
	cfg, err := config.InitConfig(cfgPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	//init logger
	log, err := logger.New(cfg.Env)

	//initial db connection
	db, err := database.New(cfg.Db)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	log.Info("successefully connect to database")

	//initial usecases
	usecase := usecase.New(db)

	//initial and run server
	server := server.New(cfg.Server.Port, usecase, log)
	server.Run(ctx)
}
