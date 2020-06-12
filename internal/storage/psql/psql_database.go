package internal

import (
	"database/sql"
	_ "github.com/lib/pq"
	logs "github.com/ernesto2108/AP_CreatyHelp/internal/logs"
)

type PostSqlClient struct {
	*sql.DB
}

func NewPSqlClient(source string) *PostSqlClient {

	db, err := sql.Open("postgres", source)

	if err != nil {
		logs.Log().Error("cannot create db: %s", err.Error())
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		logs.Log().Warn("cannot connect to postgres")
	}

	return &PostSqlClient{db}
}

func (c *PostSqlClient) viewStats() sql.DBStats {
	return c.Stats()
}