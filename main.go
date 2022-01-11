/*
 * @Description: 基于gin+gorm实现的bubble
 * @Author: lvhao
 * @Date: 2022-01-12 00:47:14
 */

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 声明全局的DB变量
var DB *gorm.DB

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/**
 * @description: 初始化DB
 * @param {*}
 * @return {error} 返回error
 */
func initMysqlDB() (err error) {
	dst := "root:lv2016.@tcp(localhost:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		fmt.Printf("connect database error. err:%v\n", err)
	}
	return
}

func main() {
	// 连接数据库
	err := initMysqlDB()
	if err != nil {
		panic(err)
	}

	// 数据库创建模型表
	DB.AutoMigrate(&Todo{})

	engin := gin.Default()

	// 加载静态文件
	engin.Static("/static", "static")

	// 加载模板文件
	engin.LoadHTMLGlob("templates/**")

	engin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// v1 路由组
	v1Group := engin.Group("v1")
	{
		// 待办事项

		// 添加
		v1Group.POST("/todo", func(c *gin.Context) {
			// 获取前端请求参数
			var todo Todo
			c.BindJSON(&todo)
			// 存储到数据库
			err := DB.Create(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})

		// 查看所有代办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []Todo
			err := DB.Find(&todoList).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})

		// 修改代办事项
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "id不存在",
				})
				return
			}

			var todo Todo
			err := DB.Where("id=?", id).First(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			}

			// fmt.Printf("%#v\n", todo)

			c.BindJSON(&todo)

			// fmt.Printf("%#v\n", todo)

			err = DB.Save(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})

		// 删除代办事项
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "id不存在",
				})
				return
			}

			err := DB.Where("id=?", id).Delete(&Todo{}).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					id: "deleted",
				})
			}
		})
	}

	engin.Run()
}
