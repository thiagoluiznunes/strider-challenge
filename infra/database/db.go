package database

import (
	"database/sql"
	"fmt"
	"sync"

	"strider-challenge/infra/config"

	_ "github.com/go-sql-driver/mysql"
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

		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
		instance, connErr = sql.Open("mysql", dsn)
		if connErr != nil {
			return
		}

		connErr = instance.Ping()
		if connErr != nil {
			return
		}

		logrus.Info("::database connection established")
	})

	return instance, connErr
}
