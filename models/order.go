package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "github.com/luqmansen/hanako/utils"
	"net/http"
	"time"
)

const(
	TIME_OFF  = 17
	TIME_BREAKFAST = 8
	TIME_LAUNCH = 12
	)

type Order struct {
	gorm.Model
	ProductID    uint `gorm:"REFERENCES products(id)" json:"product_id"`
	SellerID     uint `gorm:"REFERENCES sellers(id)";json:"seller_id"`
	BuyerID      uint `gorm:"REFERENCES users(id)" json:"buyer_id"`
	DeliveryTime string `json:"delivery_time"`
	Subscription bool   `json:"subscription"`
	TotalPrice uint32 `json:"total_price"`
	Paid         bool   `json:"paid"`
}

func (order *Order) ValidateOrder() (map[string]interface{}, bool) {

	if order.ProductID <=  0 {
		return u.Message(http.StatusBadRequest, "Must Select Item"), false
	}

	return u.Message(http.StatusOK, "All Requirement Satisfied"), true
}

func (order *Order) CreateOrder() (map[string]interface{}, bool){


	if resp, ok := order.ValidateOrder(); !ok{
		return  resp, false
	}

	product := &Product{}
	err := GetDB().Debug().Select("price, seller_id").Table("products").Where("id = ?", order.ProductID).First(product).Error
	if err != nil {
		fmt.Println(err)
		return u.Message(http.StatusNotFound, "Product Not Available"), false
	}
	price := product.Price

	if order.Subscription == true{
		order.TotalPrice = price * 30
	} else {
		order.TotalPrice = price
	}

	order.SellerID = product.SellerID

	var now = time.Now()
	if now.Hour() < TIME_BREAKFAST{
		order.DeliveryTime = "08.00"
	} else if now.Hour() <TIME_LAUNCH{
		order.DeliveryTime = "13.00"
	} else {
		order.DeliveryTime = "Tomorrow"
	}

	GetDB().Create(order)

	resp := u.Message(http.StatusOK, "Order Created")
	resp["order"] = order
	return resp, true
}






































































