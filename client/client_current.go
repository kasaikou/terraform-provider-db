package client

import (
	"context"
	"errors"

	_ "github.com/lib/pq"
)

type CurrentResponse struct {
	User    string
	DBName  string
	Version string
}

func (c *DatabaseClient) ClientIsSupported() bool {
	switch c.driver {
	case "postgres":
		return true
	default:
		return false
	}
}

func (c *DatabaseClient) Current(ctx context.Context) (crt CurrentResponse, err error) {
	switch c.driver {
	case "postgres":
		return c.CurrentPostgres(ctx)
	default:
		return crt, errors.ErrUnsupported
	}
}

func (c *DatabaseClient) CurrentPostgres(ctx context.Context) (crt CurrentResponse, err error) {
	row := c.db.QueryRowContext(ctx, `SELECT current_user, current_database(), version();`)
	if row == nil {
		return crt, errors.New("cannot get current status")
	}

	if err := row.Scan(&crt.User, &crt.DBName, &crt.Version); err != nil {
		return crt, err
	}

	return crt, nil
}
