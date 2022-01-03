package controllers

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"todolist/models"
)

const ContentTypeJson = "application/json"
const ContentTypeHtml = "text/html; UTF-8"

func TestGetTodo(t *testing.T) {
	type results struct {
		httpStatus int
		output     map[string]interface{}
	}

	tests := []struct {
		name        string
		path        string
		contentType string
		want        results
	}{
		{name: "malformed id", path: "/todos/a", contentType: ContentTypeJson,
			want: results{
				httpStatus: http.StatusBadRequest,
				output:     map[string]interface{}{"errMsg": "Invalid id"}},
		},
		{name: "not found", path: "/todos/100", contentType: ContentTypeJson,
			want: results{
				httpStatus: http.StatusNotFound,
				output:     map[string]interface{}{"errMsg": "Record not found"}},
		},
		{name: "invalid content type", path: "/todos/1", contentType: "bad content type",
			want: results{
				httpStatus: http.StatusBadRequest,
				output:     map[string]interface{}{"errMsg": "Unknown content type"}},
		},
		{name: "valid id", path: "/todos/1", contentType: ContentTypeJson,
			want: results{
				httpStatus: http.StatusOK,
				output: map[string]interface{}{
					"title":       "Pay bills",
					"id":          1.0,
					"description": "Gas, Electricity, Sewage",
					"dueDate":     "2022-01-01T09:00:00Z",
				}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			models.InitDb()
			defer models.ClearDb()
			router := InitRouter("../templates/**/*")

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tt.path, nil)
			req.Header.Set("Content-Type", tt.contentType)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.want.httpStatus, w.Code) //verify status code

			var got map[string]interface{}

			err := json.Unmarshal(w.Body.Bytes(), &got)
			if err != nil {
				t.Fatal(err)
			}
			//want is a gin.H that contains the wanted map
			if diff := cmp.Diff(tt.want.output, got); diff != "" {
				assert.True(t, false, diff)
			}
		})
	}
}

func TestGetTodos(t *testing.T) {
	type results struct {
		httpStatus int
		output     map[string]interface{}
	}

	tests := []struct {
		name        string
		path        string
		contentType string
		want        results
	}{
		{name: "all todos", path: "/todos/", contentType: ContentTypeJson,
			want: results{
				httpStatus: http.StatusOK,
				output: map[string]interface{}{
					"count": 3.0,
					"title": "All Todos",
				}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			models.InitDb()
			defer models.ClearDb()
			router := InitRouter("../templates/**/*")

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tt.path, nil)
			req.Header.Set("Content-Type", tt.contentType)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.want.httpStatus, w.Code) //verify status code

			var got map[string]interface{}

			err := json.Unmarshal(w.Body.Bytes(), &got)
			if err != nil {
				t.Fatal(err)
			}

			//only compare json map keys/values that are  in our 'want' map.
			//Ignore the rest.. it's too hard to list all fo the fields we expect in the returned todo objects
			opt := cmpopts.IgnoreMapEntries(func(key string, val interface{}) bool {
				val, exists := tt.want.output[key]
				if exists {
					return false
				}
				return true
			})

			if diff := cmp.Diff(tt.want.output, got, opt); diff != "" {

				assert.True(t, false, diff)
			}
		})
	}
}
