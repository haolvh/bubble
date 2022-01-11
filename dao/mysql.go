/*
 * @Description: dao层
 * @Author: lvhao
 * @Date: 2022-01-12 03:27:44
 */

package dao

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 声明全局的DB变量
var DB *gorm.DB

/**
 * @description: 初始化DB
 * @param {*}
 * @return {error} 返回error
 */
func InitMysqlDB() (err error) {
	dst := "root:lv2016.@tcp(localhost:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		fmt.Printf("connect database error. err:%v\n", err)
	}
	return
}
