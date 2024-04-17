package routers

import (
	"github.com/gin-gonic/gin"
	"go-list-demo/controller"
)

func SetUpRouters() *gin.Engine {

	//gin框架部分
	//建立路由器
	router := gin.Default()
	//资源路径
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "static")
	//主页
	router.GET("/", controller.IndexHandler)

	v1group := router.Group("v1")
	{
		//添加
		v1group.POST("/todo", controller.CreateATodo)
		//查看所有的待办事项
		v1group.GET("/todo", controller.GetTodoList)
		//修改
		v1group.PUT("/todo/:id", controller.UpdateATodo)
		//删除
		v1group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return router
}
