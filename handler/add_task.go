package handler

import (
	"encoding/json"
	"github.com/dasuken/budougumibon/entity"
	"github.com/dasuken/budougumibon/store"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"net/http"
)

// ここで依存性の逆転なのね
type AddTask struct {
	//Store *store.TaskStore
	DB        *sqlx.DB
	Repo      *store.Repository
	Validator *validator.Validate
}

func (at *AddTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var b struct {
		Title string `json:"title" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJson(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	// validate
	err := at.Validator.Struct(b)
	if err != nil {
		RespondJson(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	// mapping
	t := &entity.Task{
		Title:  b.Title,
		Status: entity.TaskStatusTodo,
	}

	// add
	err = at.Repo.AddTask(ctx, at.DB, t)
	if err != nil {
		RespondJson(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	// response
	rsp := struct {
		ID entity.TaskID `json:"id"`
	}{ID: t.ID}
	RespondJson(ctx, w, rsp, http.StatusOK)
}