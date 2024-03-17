package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	rep "github.com/LiaNestazi/moviesApi/pkg/repository"
	server "github.com/LiaNestazi/moviesApi/pkg/server"
	service "github.com/LiaNestazi/moviesApi/pkg/service"

	"github.com/spf13/viper"
	"github.com/joho/godotenv"
)

func main() {
	if err := initConfig(); err != nil {
        log.Fatalf("Failed to get configuration: %s", err.Error())
    }

	if err := godotenv.Load(); err != nil{
		log.Fatalf("Error loading env variables: %s", err.Error())
	}

	db, err := rep.NewPostgresDB(rep.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})

	if err != nil{
		log.Fatalf("Failed to initialize DB: %s", err.Error())
	} else{
		log.Println("Connected to DB")
	}
	
	repos := rep.NewRepository(db)
	services := service.NewService(repos)
	handlers := server.NewHandler(services)
	handlers.InitRoutes()

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port")); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
	log.Println("Movies API app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Movies API shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Fatalf("Error occured on db connection closing: %s", err.Error())
	}
	
}

func initConfig() error{
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}