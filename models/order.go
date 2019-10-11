package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	ProductID    string `gorm:"REFERENCES products(id)" json:"product_id"`
	SellerID     string `gorm:"REFERENCES sellers(id)";json:"seller_id"`
	BuyerID      string `gorm:"REFERENCES users(id)" json:"buyer_id"`
	DeliveryTime string `json:"delivery_time"`
	Subscription bool   `json:"subscription"`
	Paid         bool   `json:"paid"`
}
