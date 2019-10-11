package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	//fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Println(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Seller{}, &User{}, &Product{}, &Order{})
	//db.Model(&Product{}).AddForeignKey("seller_id", "sellers(id)", "RESTRICT", "RESTRICT")
	//db.Model(&Order{}).AddForeignKey("seller_id", "sellers(id)", "RESTRICT", "RESTRICT")
	//db.Model(&Order{}).AddForeignKey("buyer_id", "users(id)", "RESTRICT", "RESTRICT")
	//db.Model(&Order{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")
}

func GetDB() *gorm.DB {
	return db
}
