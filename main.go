package main

import (
	"fmt"
	"go-list-demo/dao"
	"go-list-demo/models"
	"go-list-demo/routers"
)

func main() {
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		fmt.Println(err)
	}
	//建立关联
	err = dao.DB.AutoMigrate(&models.Todo{})
	if err != nil {
		fmt.Println(err)
	}

	//启动！
	r := routers.SetUpRouters()
	err = r.Run(":8745")
	if err != nil {
		fmt.Println(err)
		return
	}
}
