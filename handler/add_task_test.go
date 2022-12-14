package handler

//
//import (
//	"bytes"
//	"github.com/dasuken/budougumibon/testutil"
//	"github.com/go-playground/validator/v10"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestAddTask(t *testing.T) {
//	t.Parallel()
//	type want struct {
//		status int
//		rspFile string
//	}
//	tests := map[string]struct {
//		reqFile string
//		want want
//	} {
//		"ok": {
//			reqFile: "testdata/add_task/ok_req.json.golden",
//			want: want{
//				status: http.StatusOK,
//				rspFile: "testdata/add_task/ok_rsp.json.golden",
//			},
//		},
//		"badRequest": {
//			reqFile: "testdata/add_task/bad_req.json.golden",
//			want: want{
//				status: http.StatusBadRequest,
//				rspFile: "testdata/add_task/bad_rsp.json.golden",
//			},
//		},
//	}
//
//	for n, tt := range tests {
//		tt := tt
//		t.Run(n, func(t *testing.T) {
//			t.Parallel()
//
//			w := httptest.NewRecorder()
//
//			// testutilからtとファイルを受け取る
//			r := httptest.NewRequest(
//				http.MethodPost,
//				"/tasks",
//				bytes.NewReader(testutil.LoadFile(t, tt.reqFile)),
//			)
//
//			sut := AddTask{
//				//Store: &store.TaskStore{
//				//	Tasks: map[entity.TaskID]*entity.Task{},
//				//},
//				//DB: ,
//				//Repo: ,
//				Validator: validator.New(),
//			}
//			sut.ServeHTTP(w, r)
//
//			resp := w.Result()
//			testutil.AssertResponse(t,
//				resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile))
//		})
//	}
//}
