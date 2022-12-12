package store

import (
	"context"
	"github.com/dasuken/budougumibon/clock"
	"github.com/dasuken/budougumibon/entity"
	"github.com/dasuken/budougumibon/testutil"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func prepareTasks(ctx context.Context, t *testing.T, con Execer) entity.Tasks {
	t.Helper()

	c := clock.FixedClocker{}
	wants := entity.Tasks {
		{
			Title: "want task 1", Status: "todo",
			Created: c.Now(), Modified: c.Now(),
		},
		{
			Title: "want task 2", Status: "todo",
			Created: c.Now(), Modified: c.Now(),
		},
		{
			Title: "want task 3", Status: "todo",
			Created: c.Now(), Modified: c.Now(),
		},
	}
	// id insert
	sql := `INERT INTO task(title, status, created, modified)
				VALUES(?, ?, ?, ?),(?, ?, ?, ?),(?, ?, ?, ?);
			`
	result, err := con.ExecContext(ctx, sql,
		wants[0].Title, wants[0].Status, wants[0].Created, wants[0].Modified,
		wants[1].Title, wants[1].Status, wants[1].Created, wants[1].Modified,
		wants[2].Title, wants[2].Status, wants[2].Created, wants[2].Modified,
		)
	if err != nil {
		t.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
	wants[0].ID = entity.TaskID(id)
	wants[1].ID = entity.TaskID(id + 1)
	wants[2].ID = entity.TaskID(id + 2)
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

// ちょっと考える力がたりない
func TestRepository_AddTask(t *testing.T) {
	//t.Parallel()
	//ctx := context.Background()
	//c := clock.FixedClocker{}
	//var wantID int64 = 20
	//okTask := &entity.Task{
	//	Title: "ok task",
	//	Status: "todo",
	//	Created: c.Now(),
	//	Modified: c.Now(),
	//}
	//
	//// 何のDB??
	//db, mock, err := sqlmock.New()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//t.Cleanup(func() { _ = db.Close() })
}