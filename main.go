package main

import (
	"fmt"
	"os"
	"strider-challenge/app/server"
	"strider-challenge/domain/service"
	"strider-challenge/infra/config"
	"strider-challenge/infra/database/repository"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})

	cfg, err := config.Read()
	if err != nil {
		EndAsErr("::fail to read config", err)
	}

	repoManager := repository.NewRepoManager(nil)
	svc := service.NewService(&cfg, repoManager)

	srv := server.NewServer(&cfg)
	err = srv.InitServer(svc)
	EndAsErr("::fail init server", err)
}

func EndAsErr(message string, err error) {
	if err != nil {
		logrus.Error(fmt.Sprintf("%s: %v", message, err))
		time.Sleep(time.Millisecond * 50)
		os.Exit(1)
	}
}
