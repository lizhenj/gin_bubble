package controller

import (
	"fmt"
	"gin_bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
 url --> controller --> logic --> model
 请求来 --> 控制器 --> 业务逻辑 --> 模型层的增删改查
*/

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateATodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := models.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(c *gin.Context) {
	var bb []*models.Todo
	fmt.Printf("wuwuwu %p\n", bb)
	todoList, err := models.GetTodoList()
	fmt.Printf("debug debug %p\n", todoList)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
		return
	}

	id1, _ := strconv.Atoi(id)
	todo, err := models.GetTodo(id1)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(todo)

	err = models.SaveTodo(todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
		return
	}

	id1, _ := strconv.Atoi(id)
	err := models.DeleteTodo(id1)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
