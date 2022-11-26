package handler

import (
	"encoding/json"
	"github.com/dasuken/budougumibon/entity"
	"github.com/dasuken/budougumibon/store"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type AddTask struct {
	Store *store.TaskStore
	Validator *validator.Validate
}

func (at *AddTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// context
	ctx := r.Context()

	// parse
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
		Title: b.Title,
		Status: entity.TaskStatusTodo,
		Created: time.Now(),
	}

	// add
	id, err := at.Store.Add(t)
	if err != nil {
		RespondJson(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	// response
	rsp := struct {
		ID entity.TaskID `json:"id"`
	}{ID: id}
	RespondJson(ctx, w, rsp, http.StatusOK)
}