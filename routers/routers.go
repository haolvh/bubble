/*
 * @Description: 路由分发
 * @Author: lvhao
 * @Date: 2022-01-12 04:09:26
 */
package routers

import (
	"bubble/controller"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 设置路由
 * @param
 * @return
 */
func SetupRouter() *gin.Engine {
	engine := gin.Default()

	// 加载静态文件
	engine.Static("/static", "static")

	// 加载模板文件
	engine.LoadHTMLGlob("templates/**")

	engine.GET("/", controller.IndexHanler)

	// v1 路由组
	v1Group := engine.Group("v1")
	{
		// 待办事项

		// 添加
		v1Group.POST("/todo", controller.CreateTodo)

		// 查看所有代办事项
		v1Group.GET("/todo", controller.GetTodoList)

		// 修改代办事项
		v1Group.PUT("/todo/:id", controller.UpdateTodo)

		// 删除代办事项
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}

	return engine
}
