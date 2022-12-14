package handler

import (
	"github.com/dasuken/budougumibon/entity"
	"github.com/dasuken/budougumibon/store"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type ListTask struct {
	//Store *store.TaskStore
	DB   *sqlx.DB
	Repo *store.Repository
}

// なんでEntityじゃダメなんだろうか
type task struct {
	ID     entity.TaskID     `json:"id"`
	Title  string            `json:"title"`
	Status entity.TaskStatus `json:"status"`
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks, err := lt.Repo.ListTasks(ctx, lt.DB)
	if err != nil {
		RespondJson(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
	}
	rsp := []task{}
	for _, t := range tasks {
		rsp = append(rsp, task{
			ID:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		})
	}
	/*
		task構造体を定義してmappingする処理が挟まってる。理由は分からない
	*/
	RespondJson(ctx, w, tasks, http.StatusOK)
}
