package utils

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Walk(r *mux.Router) {
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Println(t)
		return nil
	})
	http.Handle("/", r)
}
