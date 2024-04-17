package models

import "go-list-demo/dao"

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/*
	对Todo的增删改查操作全都放在这里
*/

// CreateATodo 创建一条todo
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	if err != nil {
		return err
	}
	return
}

func GetTodoList(todolist *[]Todo) (err error) {
	err = dao.DB.Find(&todolist).Error
	if err != nil {
		return err
	}
	return
}

func GetATodo(todo *Todo, id string) (err error) {
	if err = dao.DB.Where("id=?", id).First(&todo).Save(&todo).Error; err != nil {
		return err
	} else {
		return
	}
}

func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}

func DeleteATodo(todo *Todo) (err error) {
	err = dao.DB.Delete(&todo).Error
	return
}
