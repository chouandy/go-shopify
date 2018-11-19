package goshopify

import "fmt"

// TransactionAPI is an interface for interfacing with the transactions endpoints of
// the Shopify API.
// See: https://help.shopify.com/api/reference/transaction
type TransactionAPI interface {
	List(int, interface{}) ([]Transaction, error)
	Count(int, interface{}) (int, error)
	Get(int, int, interface{}) (*Transaction, error)
	Create(int, Transaction) (*Transaction, error)
}

// TransactionAPIOp handles communication with the transaction related methods of the
// Shopify API.
type TransactionAPIOp struct {
	client *Client
}

// TransactionResource represents the result from the orders/X/transactions/Y.json endpoint
type TransactionResource struct {
	Transaction *Transaction `json:"transaction"`
}

// TransactionsResource represents the result from the orders/X/transactions.json endpoint
type TransactionsResource struct {
	Transactions []Transaction `json:"transactions"`
}

// List transactions
func (s *TransactionAPIOp) List(orderID int, options interface{}) ([]Transaction, error) {
	path := fmt.Sprintf("%s/%d/transactions.json", ordersBasePath, orderID)
	resource := new(TransactionsResource)
	err := s.client.Get(path, resource, options)
	return resource.Transactions, err
}

// Count transactions
func (s *TransactionAPIOp) Count(orderID int, options interface{}) (int, error) {
	path := fmt.Sprintf("%s/%d/transactions/count.json", ordersBasePath, orderID)
	return s.client.Count(path, options)
}

// Get individual transaction
func (s *TransactionAPIOp) Get(orderID int, transactionID int, options interface{}) (*Transaction, error) {
	path := fmt.Sprintf("%s/%d/transactions/%d.json", ordersBasePath, orderID, transactionID)
	resource := new(TransactionResource)
	err := s.client.Get(path, resource, options)
	return resource.Transaction, err
}

// Create a new transaction
func (s *TransactionAPIOp) Create(orderID int, transaction Transaction) (*Transaction, error) {
	path := fmt.Sprintf("%s/%d/transactions.json", ordersBasePath, orderID)
	wrappedData := TransactionResource{Transaction: &transaction}
	resource := new(TransactionResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Transaction, err
}
