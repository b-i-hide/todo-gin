package controllers

import (
	"localhost/todo-gin/models"
	"github.com/gin-gonic/gin"
	"database/sql"
	"fmt"
	"net/http"
	"time"
	"log"
	"strconv"
)

type Todo struct {
	DB *sql.DB
}

func (s Todo) Create(c *gin.Context) {
	format := "2006-01-02"
	title := c.PostForm("title")
	due := c.PostForm("due")


	t, err := time.Parse(format, due)
	if err != nil {
		log.Fatalf("can't parse string to time. error: %s", err)
	}

	todo := models.NewTodo()

	todo.Title = title
	todo.Due = t

	err = todo.Insert(s.DB, todo)
	if err == nil {
		c.Redirect(301, "/")
	} else {
		c.JSON(404, gin.H{"error": "no user(s) into the table"})
	}
}

func (s Todo) Index(c *gin.Context) {
	todo := models.NewTodo()
	todos, err := todo.GetAll(s.DB)
	if err == nil {
		c.HTML(http.StatusOK, "todos_index.html", gin.H{
			"todos": todos,
			"title": "Root",
		})
	} else {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": err})
	}
}

func (s Todo) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	todo := models.NewTodo()
	err = todo.RemoveById(s.DB, id)
	if err == nil {
		c.Redirect(301, "/")
	} else {
		log.Fatalf("error: %s", err)
	}
}

func (s Todo) Edit(c *gin.Context) {
	format := "2006-01-02"
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	todo := models.NewTodo()
	todo, err = todo.GetById(s.DB, id)
	t := todo.Created_at
	if err == nil {
		c.HTML(http.StatusOK, "todos_edit.html", gin.H{
			"title": "Edit",
			"todo": todo,
			"created_at": t.Format(format),
		})
	} else {
		log.Fatalf("error: %s", err)
	}

}

func (s Todo) Update(c *gin.Context) {
	format := "2006-01-02"
	todo := models.NewTodo()
	title := c.PostForm("title")
	due := c.PostForm("due")
	note := c.PostForm("note")

	t, err := time.Parse(format, due)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	todo, err = todo.GetById(s.DB, id)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	todo.Title = title
	todo.Due = t
	todo.Note = note

	err = todo.Update(s.DB, todo)
	path := "/todos/" + strconv.Itoa(int(id))
	if err == nil {
		c.Redirect(301, path)
	} else {
		log.Fatalf("error: %s", err)
	}
}

func (s Todo) Show(c *gin.Context)  {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

 	todo := models.NewTodo()
 	todo, err = todo.GetById(s.DB, id)
 	if err == nil {
 		c.HTML(http.StatusOK, "todos_show.html", gin.H{
 			"title": "Show",
 			"todo": todo,
		})
	} else {
		log.Fatalf("error: %s", err)
	}
}
