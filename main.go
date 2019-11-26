package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/luqmansen/Coolinary/controllers"
	"net/http"
	"os"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,"<h1>Hello there, go try out some stuff here!</h1>")
	})


	api := r.PathPrefix("/api").Subrouter()

	//api.Use(app.JwtAuthentication)

	//PRODUCT
	api.Path("/Food").Queries("show", "{show}").HandlerFunc(controllers.GetAllProduct).Methods("GET")
	api.Path("/Food").HandlerFunc(controllers.GetAllProduct).Methods("GET")

	//USER STUFF
	api.HandleFunc("/user/new", controllers.CreateUserAccount).Methods("POST")
	api.HandleFunc("/user/login", controllers.AuthenticateUser).Methods("POST")
	api.HandleFunc("/user/order/new", controllers.CreateOrder).Methods("POST")
	api.HandleFunc("/user/order/pay/{id}", controllers.PayOrder).Methods("POST")
	api.HandleFunc("/user/order/cancel/{id}", controllers.CancelOrder).Methods("POST")
	api.HandleFunc("/user/order/skiptoday/{id}", controllers.SkipToday).Methods("POST")

	//SELLER STUFF
	api.HandleFunc("/seller/new", controllers.CreateSellerAccount).Methods("POST")
	api.HandleFunc("/seller/login", controllers.AuthenticateSeller).Methods("POST")
	//api.HandleFunc("/seller/product/new", controllers.CreateProduct).Methods("POST")
	api.HandleFunc("/Food", controllers.CreateProduct).Methods("POST")
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
