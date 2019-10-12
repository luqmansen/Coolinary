package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	u "github.com/luqmansen/Coolinary/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strings"
)

type Seller struct {
	gorm.Model
	Name         string `json:"name"`
	StoreName    string `json:"store_name"`
	StoreAddress string `json:"store_address"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Token        string `json:"token";sql:"-"`
}

func (seller *Seller) ValidateSeller() (map[string]interface{}, bool) {

	if seller.Name == "" {
		return u.Message(http.StatusBadRequest, "Name can't be empty"), false
	}

	if !strings.Contains(seller.Email, "@") {
		return u.Message(http.StatusBadRequest, "Email Address is Required"), false
	}

	if len(seller.Password) < 6 {
		return u.Message(http.StatusBadRequest, "Password is Required more than 6 character"), false
	}

	temp := &Seller{}

	err := GetDB().Table("sellers").Where("email = ?", seller.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(http.StatusInternalServerError, u.ConnectionError), false
	}
	if temp.Email != "" {
		return u.Message(http.StatusBadRequest, "Email address already in use by another seller."), false
	}

	temp2 := &Seller{}

	err2 := GetDB().Table("sellers").Where("store_name = ?", seller.StoreName).First(temp2).Error
	if err2 != nil && err2 != gorm.ErrRecordNotFound {
		return u.Message(http.StatusInternalServerError, u.ConnectionError), false
	}
	if temp2.StoreName != "" {
		return u.Message(http.StatusBadRequest, "Store Name already in use by another seller."), false
	}
	if seller.StoreAddress == "" {
		return u.Message(http.StatusBadRequest, "Address can't be empty"), false
	}

	return u.Message(http.StatusOK, "All Requirement Satisfied"), true

}

func (seller *Seller) CreateStore() map[string]interface{} {

	if resp, ok := seller.ValidateSeller(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(seller.Password), bcrypt.DefaultCost)
	seller.Password = string(hashedPassword)

	GetDB().Create(seller)

	tk := &Token{UserId: seller.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	seller.Token = tokenString

	seller.Password = ""

	response := u.Message(http.StatusOK, "New Account Created")
	response["seller"] = seller
	return response

}

func LoginSeller(email, password string) map[string]interface{} {

	seller := &Seller{}
	err := GetDB().Table("sellers").Where("email = ?", email).First(seller).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(http.StatusNotFound, u.DataNotFound)
		}
		return u.Message(http.StatusInternalServerError, u.ConnectionError)
	}

	err = bcrypt.CompareHashAndPassword([]byte(seller.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(http.StatusForbidden, u.InvalidLogin)
	}

	seller.Password = ""

	tk := &Token{UserId: seller.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	seller.Token = tokenString

	resp := u.Message(http.StatusOK, "Logged In")
	resp["seller"] = seller
	return resp

}

func GetSeller(s uint) *Seller {

	seller := &Seller{}
	GetDB().Table("sellers").Where("id = ?", s).First(seller)
	if seller.Email == "" {
		return nil
	}
	seller.Password = ""
	return seller
}