package model

import (
	"app/internal/config"
	"app/internal/log"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var _db *sqlx.DB

func Initialize() error {
	dsn := config.GetMySQLDSN()
	if dsn == "" {
		return errors.New("Empty MySQL DSN")
	}
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	_db = db
	return nil
}

func Terminate() {
	if _db == nil {
		return
	}
	if err := _db.Close(); err != nil {
		log.Logger().Error("Terminate database error: ", err)
		return
	}
}
