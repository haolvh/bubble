/*
 * @Description: todo模型
 * @Author: lvhao
 * @Date: 2022-01-12 03:31:26
 */

package entity

import "bubble/dao"

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// Todo 的增删改查都放在这里

/**
 * @description: 创建todo
 * @param {*Todo} todo
 * @return {error}
 */
func CreateTodo(todo *Todo) (err error) {
	err = dao.DB.Create(todo).Error
	if err != nil {
		return err
	}
	return
}

/**
 * @description: 获取所有的todo
 * @param
 * @return {[]*todo, error}
 */
func GetTodoList() (todoList []*Todo, err error) {
	err = dao.DB.Find(&todoList).Error
	if err != nil {
		return nil, err
	}
	return todoList, nil
}

/**
 * @description: 根据id查询todo
 * @param {string} id
 * @return {*Todo, error}
 */
func GetTodoById(id string) (todo *Todo, err error) {
	todo = new(Todo)
	err = dao.DB.Where("id=?", id).Find(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

/**
 * @description: 更新todo
 * @param {*Todo} todo
 * @return {error}
 */
func UpdateTodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

/**
 * @description: 删除todo
 * @param {string} id
 * @return {error}
 */
func DeleteTodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
