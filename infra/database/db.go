package database

import (
	"database/sql"
	"strider-challenge/infra/config"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var (
	instance     *sql.DB
	onceDataBase sync.Once
	connErr      error
)

func ConnectDataBase(cfg *config.Config) (*sql.DB, error) {

	onceDataBase.Do(func() {
		logrus.Info("::database connection initiated")

		myConfig := mysql.Config{
			User:   cfg.DBUser,
			Passwd: cfg.DBPass,
			DBName: cfg.DBName,
			Loc:    &time.Location{},
		}
		instance, connErr = sql.Open("mysql", myConfig.FormatDSN())
		if connErr != nil {
			return
		}

		logrus.Info("::database connection established")
	})

	return instance, connErr
}
