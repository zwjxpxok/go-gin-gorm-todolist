package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

var DB *gorm.DB

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() (err error) {
	//连接数据库
	dsn := "root:a3986611@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err

}

func main() {
	err := initMySQL()
	if err != nil {
		fmt.Println(err)
	}
	//建立关联
	err = DB.AutoMigrate(&Todo{})
	if err != nil {
		fmt.Println(err)
	}

	//gin框架部分
	//建立路由器
	router := gin.Default()
	//资源路径
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "static")
	//处理GET请求
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1group := router.Group("v1")
	{
		//todo
		//添加
		v1group.POST("/todo", func(context *gin.Context) {
			//前端页面点击提交待办事项 发送请求到这里
			//从请求中取出数据
			var todo Todo
			err := context.BindJSON(&todo)
			if err != nil {
				fmt.Println(err)
			}
			//存入数据库
			err = DB.Create(&todo).Error
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			}
			//给前端返回响应
			context.JSON(http.StatusOK, todo)
		})
		//查看所有的待办事项
		v1group.GET("/todo", func(context *gin.Context) {
			var todolist []Todo
			err := DB.Find(&todolist).Error
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": err,
				})
			}
			context.JSON(http.StatusOK, todolist)
		})
		//查看某一个待办事项
		v1group.GET("/todo/:id", func(context *gin.Context) {
			//1.从请求中获取要查询的id
			//2.在数据库中查询该id的数据
			//3.返回响应
		})
		//修改
		v1group.PUT("/todo/:id", func(c *gin.Context) {
			//1.获取要修改的id
			id := c.Param("id")
			//if !ok {
			//	c.JSON(http.StatusOK, gin.H{
			//		"error": "无效的id",
			//	})
			//	return
			//}
			//2.在数据库中查询该id的数据
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
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
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err,
				})
				return
			} else {
				c.JSON(http.StatusOK, todo)
			}

		})
		//删除
		v1group.DELETE("/todo/:id", func(c *gin.Context) {
			//1.从请求中获取要查询的id
			id := c.Param("id")
			//2.在数据库中删除该id的数据
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err,
				})
				return
			}

			if err = DB.Delete(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err,
				})
				return
			} else {
				c.JSON(http.StatusOK, todo)
			}
			//3.返回响应
		})
	}
	err = router.Run(":8745")
	if err != nil {
		fmt.Println(err)
		return
	}
}
