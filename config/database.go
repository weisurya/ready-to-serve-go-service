package config

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DatabaseSetting struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
}

func (dbSetting DatabaseSetting) InitiateDatabase() (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"postgres",
		"postgres://"+dbSetting.Username+
			":"+dbSetting.Password+
			"@"+dbSetting.Host+
			":"+dbSetting.Port+
			"/"+dbSetting.Name,
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
