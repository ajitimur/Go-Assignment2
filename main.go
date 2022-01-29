package main

import (
	"assignment2/controller"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var PORT = ":8080"

func main() {

	//router
	r := mux.NewRouter()

	r.HandleFunc("/orders", controller.PostOrders).Methods("POST")
	r.HandleFunc("/orders", controller.GetAllOrders).Methods("GET")
	r.HandleFunc("/orders/{id}", controller.DeleteOrder).Methods("DELETE")
	r.HandleFunc("/orders/{id}", controller.UpdateOrder).Methods("PUT")

	fmt.Println("application is running on PORT", PORT)
	http.ListenAndServe(PORT, r)
}
