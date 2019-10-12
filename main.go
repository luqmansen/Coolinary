package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/luqmansen/Coolinary/app"
	"github.com/luqmansen/Coolinary/controllers"
	"net/http"
	"os"
)

func main() {

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	api.Use(app.JwtAuthentication)

	//USER STUFF
	api.HandleFunc("/user/new", controllers.CreateUserAccount).Methods("POST")
	api.HandleFunc("/user/login", controllers.AuthenticateUser).Methods("POST")
	api.HandleFunc("/user/order/new", controllers.CreateOrder).Methods("POST")

	//SELLER STUFF
	api.HandleFunc("/seller/new", controllers.CreateSellerAccount).Methods("POST")
	api.HandleFunc("/seller/login", controllers.AuthenticateSeller).Methods("POST")
	api.HandleFunc("/seller/product/new", controllers.CreateProduct).Methods("POST")


	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Println(port)

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Println(err)
	}

}
