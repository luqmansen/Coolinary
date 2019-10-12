package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "github.com/luqmansen/hanako/utils"
	"net/http"
)

const TIME_OFF  = "05:00"

type Order struct {
	gorm.Model
	product *Product
	//ProductID    uint `gorm:"REFERENCES products(id)" json:"product_id"`
	//SellerID     uint `gorm:"REFERENCES sellers(id)";json:"seller_id"`
	BuyerID      uint `gorm:"REFERENCES users(id)" json:"buyer_id"`
	DeliveryTime string `json:"delivery_time"`
	Subscription bool   `json:"subscription"`
	Paid         bool   `json:"paid"`
}

func (order *Order) ValidateOrder() (map[string]interface{}, bool) {

	if order.product.ID <=  0 {
		return u.Message(http.StatusBadRequest, "Must Select Item"), false
	}

	return u.Message(http.StatusOK, "All Requirement Satisfied"), true
}

func (order *Order) CreateOrder(userID uint) (map[string]interface{}, bool){


	if resp, ok := order.ValidateOrder(); !ok{
		return  resp, false
	}

	temp := &Order{}
	err := GetDB().Table("products").Where("id = ?", order.product.ID).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(http.StatusInternalServerError, "Connection error, Please Retry"), false
	}

	if temp.product == nil  {
		return u.Message(http.StatusNotFound, "Product Not Available"), false
	}

	order.BuyerID = userID

	//For now, item comes from specific seller
	products := &Product{}
	err = GetDB().Table("products").Where("id = ?", order.product.ID).First(products).Error
	if err != nil {
		fmt.Println(err)
	}
	order.product = products

	if order.Subscription == true{
		order.product.Price = order.product.Price * 30
	}

	GetDB().Create(order)

	resp := u.Message(http.StatusOK, "Order Created")
	resp["order"] = order
	return resp, true



	//user_area, _ := GetDB().Table("users").Where("id = ?", userID).Get("address")



	//err := product.Where("selling_area ILIKE '%' || ? || '%'", user_area).First(product).Error
	//if err != nil  && err != gorm.ErrRecordNotFound{
	//	fmt.Println(err)
	//	return  nil
	//} else if err ==



}






































































