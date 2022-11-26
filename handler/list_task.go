package handler

import (
	"github.com/dasuken/budougumibon/store"
	"net/http"
)

type ListTask struct {
	Store *store.TaskStore
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks := lt.Store.All()
	/*
		task構造体を定義してmappingする処理が挟まってる。理由は分からない
	*/
	RespondJson(ctx, w, tasks, http.StatusOK)
}