package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"todo-api/models"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	m.Run()
}

func TestGetTodos(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/todos", nil)
	router.ServeHTTP(w, req)

	var responseTodos []models.Todo
	err := json.Unmarshal(w.Body.Bytes(), &responseTodos)
	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, models.Todos, responseTodos)
}

func TestGetTodoByID(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/todos/1", nil)
	router.ServeHTTP(w, req)

	var responseTodo models.Todo
	err := json.Unmarshal(w.Body.Bytes(), &responseTodo)
	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, models.Todos[0], responseTodo)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/todos/999", nil)
	router.ServeHTTP(w, req)

	var errorResponse gin.H
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, gin.H{"message": "todo not found"}, errorResponse)
}

func TestPostTodo(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	exampleTodo := models.Todo{
		Title:       "todo for test",
		Description: "My todo",
	}
	todoJson, _ := json.Marshal(exampleTodo)

	req, _ := http.NewRequest("POST", "/todos", strings.NewReader(string(todoJson)))
	router.ServeHTTP(w, req)

	var responseTodo models.Todo
	err := json.Unmarshal(w.Body.Bytes(), &responseTodo)
	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, exampleTodo.Title, responseTodo.Title)
	assert.Equal(t, exampleTodo.Description, responseTodo.Description)
	assert.NotEmpty(t, responseTodo.ID)
}
