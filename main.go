package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

var DB *gorm.DB

type todo struct {
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
		panic(err)
	}
	//迁移结构体
	err = DB.AutoMigrate(&todo{})
	if err != nil {
		panic(err)
	}

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

		})
		//查看所有的待办事项
		v1group.GET("/todo", func(context *gin.Context) {

		})
		//查看某一个待办事项
		v1group.GET("/todo/:id", func(context *gin.Context) {

		})
		//修改
		v1group.PUT("/todo:id", func(context *gin.Context) {

		})
		//删除
		v1group.DELETE("/todo:id", func(context *gin.Context) {

		})
	}
	err = router.Run(":8745")
	if err != nil {
		fmt.Println(err)
		return
	}
}
