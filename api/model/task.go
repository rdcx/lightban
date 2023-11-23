package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`

	List   *List `json:"list"`
	ListID uint  `json:"list_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
