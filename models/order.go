package models

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"

	// "fmt"
	"time"
)

type Orders struct {
	Order_id      int
	Customer_name string
	Ordered_at    time.Time
}

func PostOrders(customer_name, orderedAt, itemCode, description string, quantity int) (bool, error) {
	// Get a Tx for making transaction requests.

	tx, err := DB.Begin()
	if err != nil {
		log.Panic(err, "++++++")
		return false, err
	}
	defer tx.Rollback()
	var order = Orders{}

	err = tx.QueryRow("INSERT INTO orders (customer_name, ordered_at) VALUES ($1, $2) Returning *", customer_name, orderedAt).Scan(&order.Order_id, &order.Customer_name, &order.Ordered_at)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		// tx.Rollback()
		log.Panic(err, "<<<<<<")
		return false, err
	}

	_, err = tx.Exec("INSERT INTO items (item_code, description, quantity, order_id) VALUES ($1, $2, $3, $4)", itemCode, description, quantity, order.Order_id)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		// tx.Rollback()
		log.Panic(err, "------")
		return false, err
	}

	tx.Commit()

	return true, nil
}

func GetAllOrders() ([]Orders, error) {
	var result []Orders

	query := "SELECT * FROM orders"
	rows, err := DB.Query(query)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var customerName string

		var str string

		err = rows.Scan(&id, &customerName, &str)
		if err != nil {
			return result, err
		}
		t, err := time.Parse(time.RFC3339, str)

		if err != nil {
			fmt.Println(err)
		}

		order := Orders{
			Order_id:      id,
			Customer_name: customerName,
			Ordered_at:    t,
		}
		result = append(result, order)
	}

	return result, nil

}

func DeleteOrder(id int) (bool, error) {

	query := "DELETE FROM orders WHERE order_id=$1"
	_, err := DB.Exec(query, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func UpdateOrder(id int, customer_name, orderedAt, itemCode, description string, quantity, LineItemId int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		log.Panic(err, "++++++")
		return false, err
	}
	defer tx.Rollback()

	_, err = DB.Exec("UPDATE orders SET customer_name=$1, ordered_at=$2 WHERE order_id=$3", customer_name, orderedAt, id)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		// tx.Rollback()
		log.Panic(err, "<<<<<<")
		return false, err
	}

	_, err = DB.Exec("UPDATE items SET item_code=$1, description=$2, quantity=$3, order_id=$4 WHERE item_id=$5", itemCode, description, quantity, id, LineItemId)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		// tx.Rollback()
		log.Panic(err, "------")
		return false, err
	}

	tx.Commit()

	return true, nil
}
