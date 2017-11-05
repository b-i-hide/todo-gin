package main

import (
	"github.com/gin-gonic/gin"
	"localhost/todo-gin/controllers"
	"localhost/todo-gin/db"
	"database/sql"
)

type Server struct {
	db *sql.DB
}


func main() {
	db := db.InitDB()

	s := NewServer()
	s.db = db
	s.Routes()
}

func NewServer() Server {
	return Server{}
}

func (s *Server) Routes() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	todos := &controllers.Todo{DB: s.db}

	router.GET("/", todos.Index)
	router.POST("/todos", todos.Create)
	router.GET("/todos/:id", todos.Show)
	router.POST("/todos/:id/delete", todos.Delete)
	router.GET("/todos/:id/edit", todos.Edit)
	router.POST("/todos/:id/update", todos.Update)

	router.Run(":8080")
}
