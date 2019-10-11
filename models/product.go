package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	ProductName string `json:"product_name"`
	SellerID    string `json:"seller_id"`
	Price       string `json:"price"`
	SellingArea string `json:"selling_area"`
}

//GET ALL()

//GET BY NAME()
