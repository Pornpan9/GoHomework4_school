package main

import(
	"github.com/Pornpan9/test4/todo"
	"github.com/Pornpan9/test4/database"
	"github.com/gin-gonic/gin"
)

func main()  {
	
	database.InitDB()
	
	router := gin.Default()
	api := router.Group("/api")
	api.GET("/todos", todo.GetHandler)
	api.POST("/todos", todo.CreateHandler)
	api.PUT("/todos/:id", todo.UpdateHandler)
	api.DELETE("/todos/:id", todo.DeleteHandler)
	api.GET("/todos/:id", todo.GetByIDHandler)	
	
	router.Run(":1234")
}