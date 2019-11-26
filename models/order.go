package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "github.com/luqmansen/Coolinary/utils"
	"net/http"
	"time"
)

const (
	TIME_OFF       = 17
	TIME_BREAKFAST = 8
	TIME_LAUNCH    = 12
)

type Order struct {
	gorm.Model
	ProductID    uint   `gorm:"REFERENCES products(id)" json:"product_id"`
	SellerID     uint   `gorm:"REFERENCES sellers(id)";json:"seller_id"`
	BuyerID      uint   `gorm:"REFERENCES users(id)" json:"buyer_id"`
	DeliveryTime string `json:"delivery_time"`
	DeliverToday bool   `json:"deliver_today"`
	Subscription bool   `json:"subscription"`
	TotalPrice   uint32 `json:"total_price"`
	Paid         bool   `json:"paid"`
}

func (order *Order) ValidateOrder() (map[string]interface{}, bool) {

	if order.ProductID <= 0 {
		return u.Message(http.StatusBadRequest, "Must Select Item"), false
	}

	return u.Message(http.StatusOK, "All Requirement Satisfied"), true
}

func (order *Order) CreateOrder() (map[string]interface{}, bool) {

	if resp, ok := order.ValidateOrder(); !ok {
		return resp, false
	}

	product := &Product{}
	err := GetDB().Debug().Select("product_name,price, seller_id").Table("products").Where("id = ?", order.ProductID).First(product).Error
	if err != nil {
		fmt.Println(err)
		return u.Message(http.StatusNotFound, "Product Not Available"), false
	}

	user := &User{}
	err = GetDB().Debug().Select("address").Table("users").Where("id = ?", order.BuyerID).First(user).Error
	if err != nil {
		fmt.Println(err)
	}

	//If same product name available with selling area same as user address, the product changed to nearest seller
	if !GetDB().Debug().Select("seller_id,selling_area,product_name").Table("products").Where("selling_area  ILIKE '%' || ? || '%' AND product_name = ?", user.Address, product.ProductName).RecordNotFound() {
		err := GetDB().Debug().Select("price, seller_id, selling_area, product_name").Table("products").Where("selling_area = ? AND product_name = ?", user.Address, product.ProductName).First(product).Error
		if err != nil {
			fmt.Println(err)
		}
	}


	//price := product.Price
	//if order.Subscription == true {
	//	order.TotalPrice = price * 30
	//} else {
	//	order.TotalPrice = price
	//}

	order.SellerID = product.SellerID

	//Determine if delivery should be today of next day
	var now = time.Now()
	if now.Hour() < TIME_BREAKFAST {
		order.DeliveryTime = "08.00"
		order.DeliverToday = true
	} else if now.Hour() < TIME_LAUNCH {
		order.DeliveryTime = "13.00"
		order.DeliverToday = true
	} else {
		order.DeliverToday = false
		order.DeliveryTime = "08.00"
	}

	GetDB().Create(order)

	resp := u.Message(http.StatusOK, "Order Created")
	resp["order"] = order
	return resp, true
}

func (order *Order) PayOrder(orderID uint) (map[string]interface{}, bool) {

	if GetDB().Debug().Table("orders").Where("id = ?", orderID).First(order).RecordNotFound() {
		return u.Message(http.StatusNotFound, "Order Not Found"), false
	}

	err := GetDB().Debug().Select("paid").Table("orders").Where("id = ?", orderID).First(order).Error
	if err != nil {
		fmt.Println(err)
	}
	if order.Paid == true {
		return u.Message(http.StatusOK, "Order Already Paid"), true
	}

	err = GetDB().Debug().Select("paid").Table("orders").Where("id = ?", orderID).Update("paid", "true").Error
	if err != nil {
		fmt.Println(err)
	}
	order.ID = orderID
	order.SellerID = order.SellerID
	order.TotalPrice = order.TotalPrice
	order.Subscription = order.Subscription

	resp := u.Message(http.StatusOK, "Payment Success")
	resp["order"] = order
	return resp, true

}

func (order *Order) CancelOrder(orderID uint) (map[string]interface{}, bool) {

	if GetDB().Debug().Table("orders").Where("id = ?", orderID).First(order).RecordNotFound() {
		return u.Message(http.StatusNotFound, "Order Not Found"), false
	}

	err := GetDB().Debug().Select("paid").Table("orders").Where("id = ?", orderID).First(order).Error
	if err != nil {
		fmt.Println(err)
	}
	if order.Paid == true {
		return u.Message(http.StatusOK, "Order Already Paid, Can't Be Cancelled"), true
	}

	//This is soft delete, the vale "DeletedAt" will be set to current time
	err = GetDB().Debug().Table("orders").Where("id = ?", orderID).Delete(order).Error
	if err != nil {
		fmt.Println(err)
	}
	order.ID = orderID
	order.SellerID = order.SellerID
	order.TotalPrice = order.TotalPrice
	order.Subscription = order.Subscription

	now := time.Now()
	order.DeletedAt = &now

	resp := u.Message(http.StatusOK, "Order Canceled")
	resp["order"] = order
	return resp, true

}

func (order *Order) SkipToday(orderID uint) (map[string]interface{}, bool) {

	if GetDB().Debug().Table("orders").Where("id = ?", orderID).First(order).RecordNotFound() {
		return u.Message(http.StatusNotFound, "Order Not Found"), false
	}

	err := GetDB().Debug().Select("paid").Table("orders").Where("id = ?", orderID).First(order).Error
	if err != nil {
		fmt.Println(err)
	}
	if order.Paid == false {
		return u.Message(http.StatusOK, "Pay The Bill First"), true
	}

	err = GetDB().Debug().Select("deliver_today").Table("orders").Where("id = ?", orderID).Update("deliver_today", "false").Error
	if err != nil {
		fmt.Println(err)
	}
	order.ID = orderID
	order.DeliveryTime = order.DeliveryTime
	order.DeliverToday = false
	order.SellerID = order.SellerID
	order.TotalPrice = order.TotalPrice
	order.Subscription = order.Subscription
	order.DeliveryTime = order.DeliveryTime

	t := time.Now().AddDate(0, 0, +1)
	resp := u.Message(http.StatusOK, "Order will be sent at at "+t.String())
	resp["order"] = order
	return resp, true

}
