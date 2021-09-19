package config

import (
	"database/sql"
	"time"
	"ujklm23/restful_api/helper"
)

func NewDB() *sql.DB {
	dsn := "root:@tcp(localhost:3306)/daily"
	db, err := sql.Open("mysql", dsn)
	helper.PanicIfError(err)

	duration := 10 * time.Minute
	db.SetConnMaxLifetime(duration)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetMaxIdleConns(10)

	return db
}
