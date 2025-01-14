package models

type Plan struct {
	ID       uint    `gorm:"primary_key"`
	Name     string  `gorm:"size:100;not null"`
	Price    float64 `gorm:"not null"`
	Duration int     `gorm:"not null"`
}
