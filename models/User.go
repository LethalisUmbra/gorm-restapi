package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID        uint   `gorm:"type:varchar(100);primaryKey"`
	FirstName string `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Email     string `gorm:"not null;unique_index" json:"email"`
	Tasks     []Task `json:"tasks"`
}
