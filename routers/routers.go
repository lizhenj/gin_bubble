package routers

import (
	"gin_bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() (r *gin.Engine) {
	r = gin.Default()

	r.Static("./static", "static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.IndexHandler)

	//v1
	v1Group := r.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", controller.CreateATodo)
		//查看所有的代办事项
		v1Group.GET("/todo", controller.GetTodoList)
		//查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		//修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}

	return
}
