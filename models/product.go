package models

import "time"

type Product struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
	DeletedAt time.Time `gorm:"index" json:"deleted"`

	Name        string  `json:"name"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"image_url"`
	Price       float32 `json:"price"`
	Company     string  `json:"company"`
}
