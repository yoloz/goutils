package handler

import "database/sql"

type Handler interface {
	CreateTable(db *sql.DB) (result sql.Result, err error)
	BatchInsertSql(db *sql.DB, startNum int, batchNum int) (result sql.Result, err error)
}
