package models

import "time"

type User struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100;unique;not null"`
	Password  string `gorm:"size:100;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
