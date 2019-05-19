package setting

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Setting struct {
	Logger *log.Logger
	DB     *sqlx.DB
}

var setting Setting

func GetLog() *log.Logger {
	return setting.Logger
}

func SetLog(logger *log.Logger) {
	setting.Logger = logger
}

func GetDB() *sqlx.DB {
	return setting.DB
}

func SetDB(db *sqlx.DB) {
	setting.DB = db
}
