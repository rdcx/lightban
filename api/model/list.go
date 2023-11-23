package model

import "time"

type List struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `yaml:"name" json:"name"`

	ProjectID uint     `json:"project_id"`
	Project   *Project `json:"project"`

	Tasks []Task `json:"tasks"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
