package app

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/mochammadshenna/saza-aluminium/util/helper"
)

func NewDB() *sql.DB {
	var (
		host     = "localhost"
		port     = 5432
		user     = "root"
		password = "root"
		dbname   = "primaku_db"
	)

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s ", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	helper.PanicError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
