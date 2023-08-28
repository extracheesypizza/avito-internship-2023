package main

import (
	"avito-app"
	"avito-app/pkg/handler"
	"avito-app/pkg/repository"
	"avito-app/pkg/service"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/spf13/viper"
)

func main() {
	// read config file
	if err := initConfig(); err != nil {
		log.Fatalf("error occured reading the config: %s", err.Error())
	}

	// read .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error occured reading the env file: %s", err.Error())
	}

	//initialize db connection
	dbcfg := repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
	db, err := repository.NewPostgresDB(dbcfg)
	if err != nil {
		log.Fatalf("error occured initializing DB: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(avito.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured starting the server up: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("cfgs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
