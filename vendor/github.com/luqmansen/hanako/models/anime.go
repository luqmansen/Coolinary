package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "github.com/luqmansen/hanako/utils"
	"net/http"
	"strconv"
)

type Anime struct {
	gorm.Model
	//ID uint `json:"anime_id" gorm:"primary_key" `
	Name     string  `json:"name"`
	Genre    string  `json:"genre"`
	Kind     string  `json:"kind"`
	Episodes string  `json:"episodes"`
	Rating   float32 `json:"rating"`
	Members  int     `json:"members"`
}

func (anime *Anime) Validate() (map[string]interface{}, bool) {

	if anime.Name == "" {
		return u.Message(http.StatusBadRequest, "Name should be on payload"), false
	}
	if anime.Genre == "" {
		return u.Message(http.StatusBadRequest, "Genre should be on payload"), false
	}
	if anime.Kind == "" {
		return u.Message(http.StatusBadRequest, "Kind should be on payload"), false
	}

	return u.Message(http.StatusOK, "success"), true
}

func (anime *Anime) AddEntry() map[string]interface{} {

	if resp, ok := anime.Validate(); !ok {
		return resp
	}

	getDB().Create(anime)

	resp := u.Message(http.StatusOK, "success")
	resp["anime"] = anime
	return resp
}

func GetAll(number string) []*Anime {

	if number == "" {
		number = "20"
	} else if _, err := strconv.Atoi(number); err != nil {
		number = "20"
	}
	animes := make([]*Anime, 0)
	err := getDB().Limit(number).Find(&animes).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return animes
}

func GetByTitle(title string) []*Anime {

	animes := make([]*Anime, 0)
	err := getDB().Table("animes").Where("name ILIKE '%' || ? || '%'", title).Find(&animes).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return animes
}

func GetByID(ID uint) *Anime {

	animes := &Anime{}
	err := getDB().Table("animes").Where("ID = ?", ID).Find(&animes).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return animes
}
