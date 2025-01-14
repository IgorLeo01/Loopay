package models

import "time"

type Invoice struct {
	ID        uint    `gorm:"primary_key"`
	UserID    uint    `gorm:"not null"`
	PlanID    uint    `gorm:"not null"`
	Amount    float64 `gorm:"not null"`
	Paid      bool    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
