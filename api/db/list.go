package db

import "lightban/api/model"

func (db *DB) GetList(id uint) (*model.List, error) {
	var l model.List
	err := db.DB.Where("id = ?", id).Preload("Project").First(&l).Error

	if err != nil {
		return nil, err
	}

	return &l, nil
}

func (db *DB) GetListsByProjectID(id uint) ([]model.List, error) {
	var l []model.List
	err := db.DB.Where("project_id = ?", id).Find(&l).Error

	if err != nil {
		return nil, err
	}

	return l, nil
}

func (db *DB) DeleteListByID(id uint) error {
	return db.DB.Where("id = ?", id).Delete(&model.List{}).Error
}

func (db *DB) UpdateList(list *model.List) error {
	return db.DB.Save(list).Error
}

func (db *DB) CreateList(list *model.List) error {
	return db.DB.Create(list).Error
}
