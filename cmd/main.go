package main

import (
	"garyshker"
	"garyshker/adapters"
	"garyshker/pkg/handler"
	"garyshker/pkg/repository"
	"garyshker/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	config := adapters.ParseConfig()
	db, err := adapters.NewPostgresDB(adapters.Config{
		Host:     config.DataBaseHost,
		Port:     config.DataBasePort,
		Username: config.DataBaseUsername,
		Password: os.Getenv("DB_PASS"),
		DBName:   config.DataBaseDbname,
		SSLMode:  config.DataBaseSslmode,
	})
	if err != nil {
		logrus.Fatal("Error connecting to the database: ", err)
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(garyshker.Server)
	if err := srv.Run(config.Port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}
