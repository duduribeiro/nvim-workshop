package main

import (
	"github.com/gin-gonic/gin"
	"todo-api/controllers"
)

func main() {
	router := setupRouter()
	router.Run("localhost:8080")
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/todos", controllers.GetTodos)
	router.GET("/todos/:id", controllers.GetTodoByID)
	router.POST("/todos", controllers.PostTodos)

	return router
}
