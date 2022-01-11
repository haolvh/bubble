/*
 * @Description: 基于gin+gorm实现的bubble
 * @Author: lvhao
 * @Date: 2022-01-12 00:47:14
 */

package main

import (
	"bubble/dao"
	"bubble/entity"
	"bubble/routers"
)

func main() {
	// 连接数据库
	err := dao.InitMysqlDB()
	if err != nil {
		panic(err)
	}

	// 数据库创建模型表
	dao.DB.AutoMigrate(&entity.Todo{})

	// 注册路由
	engine := routers.SetupRouter()

	engine.Run()
}
