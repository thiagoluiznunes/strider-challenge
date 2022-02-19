package database

import (
	"database/sql"
	"net"
	"sync"
	"time"

	"github.com/thiagoluiznunes/strider-challenge/infra/config"

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

		dbAddr := net.JoinHostPort(cfg.DBHost, cfg.DBPort)
		myConfig := mysql.Config{
			Addr:                 dbAddr,
			User:                 cfg.DBUser,
			Passwd:               cfg.DBPass,
			DBName:               cfg.DBName,
			Loc:                  &time.Location{},
			ParseTime:            true,
			AllowNativePasswords: true,
		}

		instance, connErr = sql.Open("mysql", myConfig.FormatDSN())
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
