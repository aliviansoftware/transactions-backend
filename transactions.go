package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// The person Type (more like an object)
type Transaction struct {
	ID      string   `json:"id,omitempty"`
	Payee   string   `json:"payee,omitempty"`
	Date    string   `json:"date,omitempty"`
	Payment *Payment `json:"payment,omitempty"`
}
type Payment struct {
	Amount string `json:"amount,omitempty"`
	Status string `json:"status,omitempty"`
}

var transactions []Transaction

// Display all from the people var
func Transactions(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(transactions)
}

// Display a single data
func GetTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range transactions {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Transaction{})
}

// create a new item
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var transaction Transaction
	_ = json.NewDecoder(r.Body).Decode(&transaction)
	transaction.ID = params["id"]
	transactions = append(transactions, transaction)
	json.NewEncoder(w).Encode(transactions)
}

// main function to boot up everything
func main() {
	router := mux.NewRouter()
	transactions = append(transactions, Transaction{ID: "1", Payee: "John", Date: "05/09/2019", Payment: &Payment{Amount: "£100.00", Status: "Processed"}})
	transactions = append(transactions, Transaction{ID: "2", Payee: "John", Date: "05/09/2019", Payment: &Payment{Amount: "£200.00", Status: "Processed"}})
	router.HandleFunc("/transactions", Transactions).Methods("GET")
	router.HandleFunc("/transactions/{id}", GetTransaction).Methods("GET")
	router.HandleFunc("/transactions/{id}", CreateTransaction).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
