package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	u "github.com/luqmansen/hanako/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strings"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

func (account *Account) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return u.Message(http.StatusBadRequest, "Email Address is Required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(http.StatusBadRequest, "Password is Required"), false
	}

	temp := &Account{}

	//	Check for duplicate email
	err := getDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(http.StatusInternalServerError, "Connection error, Please Retry"), false
	}

	if temp.Email != "" {
		return u.Message(http.StatusConflict, "Email address already in use by another user."), false
	}

	return u.Message(http.StatusOK, "Requirement passed"), true

}

func (account *Account) Create() map[string]interface{} {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	getDB().Create(account)

	if account.ID <= 0 {
		return u.Message(http.StatusInternalServerError, "Failed To Create Account, connection Error.")
	}

	//	Crete new jwt Token for newly created account
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //Delete pass from memory

	response := u.Message(http.StatusOK, "Account Has Been Created")
	response["account"] = account
	return response

}

func Login(email, password string) map[string]interface{} {

	account := &Account{}
	err := getDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(http.StatusNotFound, "Email Address Not Found")
		}
		return u.Message(http.StatusInternalServerError, "Connection Error, Please Retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(http.StatusForbidden, "Invalid Login Credential")
	}

	//password work, and delete password immediately
	account.Password = ""

	//Create Jwt Token
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	resp := u.Message(http.StatusOK, "Logged In")
	resp["account"] = account
	return resp
}

func GetUser(u uint) *Account {

	acc := &Account{}
	getDB().Table("account").Where("id = ?", u).First(acc)
	if acc.Email == "" {
		return nil
	}
	acc.Password = ""
	return acc
}
