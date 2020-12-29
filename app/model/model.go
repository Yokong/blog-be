package model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yokowu/blog-db/db"
)

type Store interface {
	db.Querier
}

type MySQLStore struct {
	db *sql.DB
	*db.Queries
}

func NewStore(d *sql.DB) Store {
	return MySQLStore{
		db:      d,
		Queries: db.New(d),
	}
}
