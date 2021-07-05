package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderNo  string  `gorm:"index:idx_order_no" json:"order_no"`
	UserName string  `gorm:"index:idx_order_username" json:"user_name"`
	Amount   float64 `json:"amount"`
	Status   bool    `json:"status"`
	FileUrl  string  `json:"file_url"`
}
