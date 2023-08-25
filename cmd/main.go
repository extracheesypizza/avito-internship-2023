package main

import (
	"avito-app"
	"avito-app/pkg/handler"
	"avito-app/pkg/repository"
	"avito-app/pkg/service"
	"log"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error occured reading the config: %s", err.Error())
	}

	repos := repository.NewRepository()
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
