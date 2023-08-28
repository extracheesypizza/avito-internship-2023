package main

import (
	"avito-app"
	"avito-app/pkg/handler"
	"avito-app/pkg/repository"
	"avito-app/pkg/service"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func main() {
	// read config file
	if err := initConfig(); err != nil {
		logrus.Fatalf("error occured reading the config: %s", err.Error())
	}

	// read .env file
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error occured reading the env file: %s", err.Error())
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
		logrus.Fatalf("error occured initializing DB: %s", err.Error())
	}

	schemaUp :=
		`DROP TABLE IF EXISTS user_segments;
	DROP TABLE IF EXISTS operations;
	DROP TABLE IF EXISTS segments;
	
	CREATE TABLE segments (
		seg_name varchar(255)  NOT NULL,
		seg_id serial PRIMARY KEY
	);
	
	CREATE TABLE user_segments (
		user_id int  NOT NULL,
		seg_id int  REFERENCES segments (seg_id) ON DELETE CASCADE
	);
	
	CREATE TABLE operations (
		operation_id serial PRIMARY KEY,
		user_id int  NOT NULL,
		seg_id int  NOT NULL,
		operation varchar(3)  NOT NULL,
		at_timestamp timestamp  NOT NULL,
		TTL int  NOT NULL
	);`

	db.Exec(schemaUp)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(avito.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured starting the server up: %s", err.Error())
		}
	}()
	logrus.Print("App Launched")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App Closing Down...")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server closing down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on server closing down: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("cfgs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
