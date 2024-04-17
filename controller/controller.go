package controller

import (
	"github.com/gin-gonic/gin"
	"go-list-demo/models"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func CreateATodo(context *gin.Context) {
	//从请求中取出数据
	var todo models.Todo
	err := context.BindJSON(&todo)
	if err != nil {
		return
	}
	//存入数据库
	err = models.CreateATodo(&todo)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, todo)
	}

}

func GetTodoList(context *gin.Context) {
	var todolist []models.Todo
	err := models.GetTodoList(&todolist)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		context.JSON(http.StatusOK, todolist)
	}

}

func UpdateATodo(c *gin.Context) {
	//1.获取要修改的id
	id := c.Param("id")
	//2.在数据库中查询该id的数据
	var todo models.Todo
	if err := models.GetATodo(&todo, id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	}
	//3.修改该条数据的status为true
	err := c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	}
	todo.Status = true
	if err = models.UpdateATodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
func DeleteATodo(c *gin.Context) {
	//1.从请求中获取要查询的id
	id := c.Param("id")
	//2.在数据库中删除该id的数据
	var todo models.Todo
	if err := models.GetATodo(&todo, id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	}

	if err := models.DeleteATodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	} else {
		c.JSON(http.StatusOK, todo)
	}
	//3.返回响应
}
