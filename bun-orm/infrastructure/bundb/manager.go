package bundb

import (
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"os"
	"sync"
)

var (
	dbConnOnce sync.Once
	conn       *bun.DB
)

func Db() *bun.DB {
	dbConnOnce.Do(func() {
		dsn := postgresqlDsn()
		hsqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
		conn = bun.NewDB(hsqldb, pgdialect.New())
	})
	return conn
}

func postgresqlDsn() string {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	return dsn
}
