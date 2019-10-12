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

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Token    string `json:"token";sql:"-"`
}

func (user *User) ValidateUser() (map[string]interface{}, bool) {

	if user.Name == "" {
		return u.Message(http.StatusBadRequest, "Name can't be empty"), false
	}

	if !strings.Contains(user.Email, "@") {
		return u.Message(http.StatusBadRequest, "Email Address is Required"), false
	}

	temp := &User{}

	err := GetDB().Table("users").Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(http.StatusInternalServerError, u.ConnectionError), false
	}
	if temp.Email != "" {
		return u.Message(http.StatusBadRequest, "Email address already in use by another user."), false
	}

	if len(user.Password) < 6 {
		return u.Message(http.StatusBadRequest, "Password is Required more than 6 character"), false
	}

	if user.Address == "" {
		return u.Message(http.StatusBadRequest, "Address can't be empty"), false
	}

	return u.Message(http.StatusOK, "All Requirement Satisfied"), true

}

func (user *User) CreateUser() map[string]interface{} {

	if resp, ok := user.ValidateUser(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	GetDB().Create(user)

	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString

	user.Password = ""

	response := u.Message(http.StatusOK, "Account Created")
	response["user"] = user
	return response

}

func LoginUser(email, password string) map[string]interface{} {

	user := &User{}
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(http.StatusNotFound, u.DataNotFound)
		}
		return u.Message(http.StatusInternalServerError, u.ConnectionError)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(http.StatusForbidden, "Invalid Login Credential")
	}

	user.Password = ""

	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString

	resp := u.Message(http.StatusOK, "Logged In")
	resp["account"] = user
	return resp

}

func GetUser(u uint) *User {

	user := &User{}
	GetDB().Table("users").Where("id = ?", u).First(user)
	if user.Email == "" {
		return nil
	}
	user.Password = ""
	return user
}
