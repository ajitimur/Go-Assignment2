package controller

import (
	"assignment2/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Messages struct {
	Status bool
}

type CreateOrder struct {
	OrderedAt    string       `json:"orderedAt"`
	CustomerName string       `json:"customerName"`
	Items        []OrderItems `json:"items"`
}

type OrderItems struct {
	LineItemId  int    `json:"LineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func PostOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body CreateOrder

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(reqBody, &body)
	if err != nil {
		panic(err)
	}

	var res bool
	res, err = models.PostOrders(body.CustomerName, body.OrderedAt, body.Items[0].ItemCode, body.Items[0].Description, body.Items[0].Quantity)
	if err != nil {
		panic(err)
	}

	var message Messages
	message.Status = res

	json.NewEncoder(w).Encode(message)
}

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, err := models.GetAllOrders()
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)
	idString := param["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	res, err := models.DeleteOrder(id)
	if err != nil {
		panic(err)
	}

	var message Messages
	message.Status = res

	json.NewEncoder(w).Encode(message)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)
	idString := param["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	var body CreateOrder

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(reqBody, &body)
	if err != nil {
		panic(err)
	}

	res, err := models.UpdateOrder(id, body.CustomerName, body.OrderedAt, body.Items[0].ItemCode, body.Items[0].Description, body.Items[0].Quantity, body.Items[0].LineItemId)
	if err != nil {
		panic(err)
	}

	var message Messages
	message.Status = res

	json.NewEncoder(w).Encode(message)
}
