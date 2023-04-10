package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID    int  `json:"id,omitempty" gorm:"primaryKey"`
	Name  string `json:"name"`
	SKU   string `json:"sku"`
	Stock int  `json:"stock"`
	Price int  `json:"price"`
}
