package main

import (
	"fmt"
	"os"
	"time"

	"github.com/thiagoluiznunes/strider-challenge/app/server"
	"github.com/thiagoluiznunes/strider-challenge/domain/service"
	"github.com/thiagoluiznunes/strider-challenge/infra/config"
	"github.com/thiagoluiznunes/strider-challenge/infra/database"
	"github.com/thiagoluiznunes/strider-challenge/infra/database/repository"

	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})

	cfg, err := config.Read()
	EndAsErr("::fail to read config", err)

	dbConn, err := database.ConnectDataBase(&cfg)
	EndAsErr("::fail connect to database", err)

	repoManager := repository.NewRepoManager(dbConn)
	svc := service.NewService(&cfg, repoManager)

	srv := server.NewServer(&cfg)
	err = srv.InitServer(svc)
	EndAsErr("::fail to init server", err)
}

func EndAsErr(message string, err error) {
	if err != nil {
		logrus.Error(fmt.Sprintf("%s: %v", message, err))
		time.Sleep(time.Millisecond * 50)
		os.Exit(1)
	}
}
