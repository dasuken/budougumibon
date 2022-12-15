package testutil

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func OpenDBForTest(t *testing.T) *sqlx.DB {
	t.Helper()

	port := 55432
	if _, defined := os.LookupEnv("CI"); defined {
		port = 5432
	}
	fmt.Println(port)
	db, err := sql.Open("postgres",
		fmt.Sprintf(
			//"todo:todo@tcp(127.0.0.1:%d)/todo?parseTime=true",
			"postgres://todo:todo@127.0.0.1:%d/todo?sslmode=disable",
			port,
		))
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		t.Fatalf("ping error: port: %d, err: %s", port, err.Error())
	}
	t.Cleanup(
		func() { _ = db.Close() },
	)
	return sqlx.NewDb(db, "postgres")
}
