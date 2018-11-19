package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const customersBasePath = "admin/customers"
const customersResourceName = "customers"

// CustomerAPI is an interface for interfacing with the customers endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/customer
type CustomerAPI interface {
	List(interface{}) ([]Customer, error)
	Count(interface{}) (int, error)
	Get(int, interface{}) (*Customer, error)
	Search(interface{}) ([]Customer, error)
	Create(Customer) (*Customer, error)
	Update(Customer) (*Customer, error)
	Delete(int) error
	ListOrders(int, interface{}) ([]Order, error)

	// MetafieldsAPI used for Customer resource to communicate with Metafields resource
	MetafieldsAPI
}

// CustomerAPIOp handles communication with the product related methods of
// the Shopify API.
type CustomerAPIOp struct {
	client *Client
}

// Customer represents a Shopify customer
type Customer struct {
	ID                  int                `json:"id,omitempty"`
	Email               string             `json:"email,omitempty"`
	FirstName           string             `json:"first_name,omitempty"`
	LastName            string             `json:"last_name,omitempty"`
	State               string             `json:"state,omitempty"`
	Note                string             `json:"note,omitempty"`
	VerifiedEmail       bool               `json:"verified_email,omitempty"`
	MultipassIdentifier string             `json:"multipass_identifier,omitempty"`
	OrdersCount         int                `json:"orders_count,omitempty"`
	TaxExempt           bool               `json:"tax_exempt,omitempty"`
	TotalSpent          *decimal.Decimal   `json:"total_spent,omitempty"`
	Phone               string             `json:"phone,omitempty"`
	Tags                string             `json:"tags,omitempty"`
	LastOrderID         int                `json:"last_order_id,omitempty"`
	LastOrderName       string             `json:"last_order_name,omitempty"`
	AcceptsMarketing    bool               `json:"accepts_marketing,omitempty"`
	DefaultAddress      *CustomerAddress   `json:"default_address,omitempty"`
	Addresses           []*CustomerAddress `json:"addresses,omitempty"`
	CreatedAt           *time.Time         `json:"created_at,omitempty"`
	UpdatedAt           *time.Time         `json:"updated_at,omitempty"`
	Metafields          []Metafield        `json:"metafields,omitempty"`
}

// CustomerResource represents the result from the customers/X.json endpoint
type CustomerResource struct {
	Customer *Customer `json:"customer"`
}

// CustomersResource represents the result from the customers.json endpoint
type CustomersResource struct {
	Customers []Customer `json:"customers"`
}

// CustomerSearchOptions represents the options available when searching for a customer
type CustomerSearchOptions struct {
	Page   int    `url:"page,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Fields string `url:"fields,omitempty"`
	Order  string `url:"order,omitempty"`
	Query  string `url:"query,omitempty"`
}

// List customers
func (s *CustomerAPIOp) List(options interface{}) ([]Customer, error) {
	path := fmt.Sprintf("%s.json", customersBasePath)
	resource := new(CustomersResource)
	err := s.client.Get(path, resource, options)
	return resource.Customers, err
}

// Count customers
func (s *CustomerAPIOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", customersBasePath)
	return s.client.Count(path, options)
}

// Get customer
func (s *CustomerAPIOp) Get(customerID int, options interface{}) (*Customer, error) {
	path := fmt.Sprintf("%s/%v.json", customersBasePath, customerID)
	resource := new(CustomerResource)
	err := s.client.Get(path, resource, options)
	return resource.Customer, err
}

// Create a new customer
func (s *CustomerAPIOp) Create(customer Customer) (*Customer, error) {
	path := fmt.Sprintf("%s.json", customersBasePath)
	wrappedData := CustomerResource{Customer: &customer}
	resource := new(CustomerResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Customer, err
}

// Update an existing customer
func (s *CustomerAPIOp) Update(customer Customer) (*Customer, error) {
	path := fmt.Sprintf("%s/%d.json", customersBasePath, customer.ID)
	wrappedData := CustomerResource{Customer: &customer}
	resource := new(CustomerResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Customer, err
}

// Delete an existing customer
func (s *CustomerAPIOp) Delete(customerID int) error {
	path := fmt.Sprintf("%s/%d.json", customersBasePath, customerID)
	return s.client.Delete(path)
}

// Search customers
func (s *CustomerAPIOp) Search(options interface{}) ([]Customer, error) {
	path := fmt.Sprintf("%s/search.json", customersBasePath)
	resource := new(CustomersResource)
	err := s.client.Get(path, resource, options)
	return resource.Customers, err
}

// ListMetafields list metafields for a customer
func (s *CustomerAPIOp) ListMetafields(customerID int, options interface{}) ([]Metafield, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: customersResourceName, resourceID: customerID}
	return metafieldAPI.List(options)
}

// CountMetafields count metafields for a customer
func (s *CustomerAPIOp) CountMetafields(customerID int, options interface{}) (int, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: customersResourceName, resourceID: customerID}
	return metafieldAPI.Count(options)
}

// GetMetafield get individual metafield for a customer
func (s *CustomerAPIOp) GetMetafield(customerID int, metafieldID int, options interface{}) (*Metafield, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: customersResourceName, resourceID: customerID}
	return metafieldAPI.Get(metafieldID, options)
}

// CreateMetafield create a new metafield for a customer
func (s *CustomerAPIOp) CreateMetafield(customerID int, metafield Metafield) (*Metafield, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: customersResourceName, resourceID: customerID}
	return metafieldAPI.Create(metafield)
}

// UpdateMetafield update an existing metafield for a customer
func (s *CustomerAPIOp) UpdateMetafield(customerID int, metafield Metafield) (*Metafield, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: customersResourceName, resourceID: customerID}
	return metafieldAPI.Update(metafield)
}

// DeleteMetafield delete an existing metafield for a customer
func (s *CustomerAPIOp) DeleteMetafield(customerID int, metafieldID int) error {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: customersResourceName, resourceID: customerID}
	return metafieldAPI.Delete(metafieldID)
}

// ListOrders retrieves all orders from a customer
func (s *CustomerAPIOp) ListOrders(customerID int, options interface{}) ([]Order, error) {
	path := fmt.Sprintf("%s/%d/orders.json", customersBasePath, customerID)
	resource := new(OrdersResource)
	err := s.client.Get(path, resource, options)
	return resource.Orders, err
}
