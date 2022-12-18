package store

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dasuken/budougumibon/clock"
	"github.com/dasuken/budougumibon/entity"
	"github.com/dasuken/budougumibon/testutil"
	"github.com/google/go-cmp/cmp"
	"github.com/jmoiron/sqlx"
	"regexp"
	"testing"
)

func prepareTasks(ctx context.Context, t *testing.T, con Execer) entity.Tasks {
	t.Helper()

	c := clock.FixedClocker{}
	wants := entity.Tasks{
		{
			ID: 1, Title: "want task 1", Status: "todo",
			Created: c.Now(), Modified: c.Now(),
		},
		{
			ID: 2, Title: "want task 2", Status: "todo",
			Created: c.Now(), Modified: c.Now(),
		},
		{
			ID: 3, Title: "want task 3", Status: "todo",
			Created: c.Now(), Modified: c.Now(),
		},
	}

	sql := `INSERT INTO task(id, title, status, created, modified)
				VALUES(:id, :title, :status, :created, :modified)
				RETURNING id
			`

	dataMap := []map[string]interface{}{
		{"id": wants[0].ID, "title": wants[0].Title, "status": wants[0].Status, "created": wants[0].Created, "modified": wants[0].Modified},
		{"id": wants[1].ID, "title": wants[1].Title, "status": wants[1].Status, "created": wants[1].Created, "modified": wants[1].Modified},
		{"id": wants[2].ID, "title": wants[2].Title, "status": wants[2].Status, "created": wants[2].Created, "modified": wants[2].Modified},
	}

	result, err := con.NamedExecContext(ctx, sql, dataMap)
	fmt.Println(result)
	if err != nil {
		t.Fatal(err)
	}
	return wants
}

func TestRepository_ListTasks(t *testing.T) {
	ctx := context.Background()
	tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)
	t.Cleanup(func() { _ = tx.Rollback() })
	if err != nil {
		t.Fatal(err)
	}

	wants := prepareTasks(ctx, t, tx)

	sut := &Repository{}
	gots, err := sut.ListTasks(ctx, tx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if d := cmp.Diff(gots, wants); len(d) != 0 {
		t.Errorf("differs: (-git +want)\n%s", d)
	}
}

func TestRepository_AddTask(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	c := clock.FixedClocker{}

	var wantID int64 = 20
	okTask := &entity.Task{
		Title:    "ok task",
		Status:   "todo",
		Created:  c.Now(),
		Modified: c.Now(),
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { _ = db.Close() })
	mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO task (title, status, created, modified) VALUES (?,?,?,?)"),
	).WithArgs(okTask.Title, okTask.Status, okTask.Created, okTask.Modified).
		WillReturnResult(sqlmock.NewResult(wantID, 1))

	// sqlmockで生成したconnectionを内包するためにsqlx.Open使わないんだわ
	xdb := sqlx.NewDb(db, "postgres")

	r := &Repository{Clocker: c}
	if err := r.AddTask(ctx, xdb, okTask); err != nil {
		t.Errorf("want no error, but got %v", err)
	}
}
