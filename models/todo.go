package models

import (
	"gin_bubble/dao"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

//Todo 增删改查

func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(todo).Error
	return
}

func GetTodoList() (todoList []*Todo, err error) {
	//fmt.Printf("debug debug %v", todoList)
	err = dao.DB.Find(&todoList).Error
	//fmt.Printf("debug debug %v", todoList)
	return
}

func GetTodo(id int) (todo *Todo, err error) {
	todo = new(Todo)
	err = dao.DB.Where("id=?", id).First(todo).Error
	return
}

func SaveTodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteTodo(id int) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
