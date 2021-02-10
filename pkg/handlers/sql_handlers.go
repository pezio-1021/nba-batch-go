package handlers

import (

	"github.com/jmoiron/sqlx"
	"https://github.com/shohei1021/nba-batch-go/pkg/config"
)

type SQLHandler struct {
	Conn *sqlx.DB
}

func NewSQLHandler(db config.DBConfig) (*SQLHandler, error) {
	conn, err := sqlx.Open("nrmysql", db.FormatDSN())
	if err != nil {
		return &SQLHandler{}, err	
	}

	if err := conn.Ping(); err != nil {
		return &SQLHandler{}, err	
	}

	return &SQLHandler{
		Conn: conn,
	}, nil
}

func (handler *SQLHandler) Get(ctx content.Context, dest interface{}, query string, args ...interface) error {
	return handler.Conn.GetContext(ctx, dest, query, args...)
}