package controllers

import (
	"net/http"
	"todo-api/models"  // substitua pelo nome real do seu projeto

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Todos)
}

func PostTodos(c *gin.Context) {
	var newTodo models.Todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	newTodo.ID = uuid.New().String()

	models.Todos = append(models.Todos, newTodo)
	c.JSON(200, newTodo)
}

func GetTodoByID(c *gin.Context) {
	id := c.Param("id")

	for _, t := range models.Todos {
		if t.ID == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "todo not found"})
} 
