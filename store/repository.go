package store

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dasuken/budougumibon/config"
	"github.com/jmoiron/sqlx"
	"time"

	_ "github.com/lib/pq"
)

func New(ctx context.Context, cfg *config.Config) (*sqlx.DB, func(), error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
			))
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, func() {_ = db.Close() }, err
	}

	xdb := sqlx.NewDb(db, "postgres")
	return xdb, func() { _ = db.Close() }, nil
}