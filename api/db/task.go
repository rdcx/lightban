package db

import "lightban/api/model"

func (db *DB) GetTask(id uint) (*model.Task, error) {
	var c model.Task
	err := db.DB.Where("id = ?", id).Preload("List.Project").First(&c).Error

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (db *DB) UpdateTask(task *model.Task) error {
	return db.DB.Save(task).Error
}

func (db *DB) DeleteTaskByID(id uint) error {
	return db.DB.Where("id = ?", id).Delete(&model.Task{}).Error
}

func (db *DB) CreateTask(task *model.Task) error {
	return db.DB.Create(task).Error
}
