/*
 * @Description: controller层
 * @Author: lvhao
 * @Date: 2022-01-12 03:15:15
 */

package controller

import (
	"bubble/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 显示index.html
 * @param {*gin.Context} c
 * @return
 */
func IndexHanler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

/**
 * @description: 新增todo
 * @param {*gin.Context} c
 * @return
 */
func CreateTodo(c *gin.Context) {
	// 获取前端请求参数
	var todo entity.Todo
	c.BindJSON(&todo)
	// 存储到数据库
	err := entity.CreateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

/**
 * @description: 获取todoList
 * @param {*gin.Context} c
 * @return
 */
func GetTodoList(c *gin.Context) {
	todoList, err := entity.GetTodoList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

/**
 * @description: 更新todo
 * @param {*gin.Context} c
 * @return
 */
func UpdateTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "id不存在",
		})
		return
	}

	todo, err := entity.GetTodoById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}

	// fmt.Printf("%#v\n", todo)

	c.BindJSON(todo)

	// fmt.Printf("%#v\n", todo)

	err = entity.UpdateTodo(todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

/**
 * @description: 删除todo
 * @param {*gin.Context} c
 * @return
 */
func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "id不存在",
		})
		return
	}

	err := entity.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			id: "deleted",
		})
	}
}
