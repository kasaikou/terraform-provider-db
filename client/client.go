package client

import "database/sql"

type DatabaseClient struct {
	driver string
	db     *sql.DB
}

func New(driver string, dsn string) (*DatabaseClient, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	return &DatabaseClient{driver: driver, db: db}, db.Ping()
}
