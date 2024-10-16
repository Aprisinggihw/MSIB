package models

import (
	"tugas-4/config"
	"tugas-4/entities"
)

func CreateTodo(todo *entities.Todo) error {
	return config.DB.Create(todo).Error
}


func GetTodos() ([]entities.Todo, error) {
	var todos []entities.Todo
	err := config.DB.Find(&todos).Error
	return todos, err
}
// Find Todo by ID
func FindTodoByID(id uint) (*entities.Todo, error) {
	var todo entities.Todo
	err := config.DB.First(&todo, id).Error
	return &todo, err
}
// Update Todo by ID
func UpdateTodo(todo *entities.Todo) error {
	return config.DB.Save(todo).Error
}

// Delete Todo by ID
func DeleteTodo(id uint) error {
	return config.DB.Delete(&entities.Todo{}, id).Error
}