package db

import (
	"strider-challenge/infra/config"
	"sync"

	"github.com/sirupsen/logrus"
)

func ConnectDataBase(cfg *config.Config) (conn interface{}, connErr error) {

	var onceDataBase sync.Once

	onceDataBase.Do(func() {
		logrus.Info("intiate connection to database")

		// TODO: create connection
		conn = nil
		connErr = nil

		logrus.Info("connected to DataBase")
	})

	return conn, connErr
}
