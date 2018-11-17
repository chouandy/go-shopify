package goshopify

import (
	"fmt"
)

const fulfillmentServicesBasePath = "admin/fulfillment_services"

// FulfillmentServiceAPI is an interface for interfacing with the fulfillment service
// endpoints of the Shopify API.
// https://help.shopify.com/en/api/reference/shipping_and_fulfillment/fulfillmentservice
type FulfillmentServiceAPI interface {
	List(interface{}) ([]FulfillmentService, error)
	Get(int, interface{}) (*FulfillmentService, error)
	Create(FulfillmentService) (*FulfillmentService, error)
	Update(FulfillmentService) (*FulfillmentService, error)
	Delete(int) error
}

// FulfillmentService fulfillment service struct
type FulfillmentService struct {
	ID                     int    `json:"id"`
	Name                   string `json:"name"`
	Handle                 string `json:"handle"`
	IncludePendingStock    *bool  `json:"include_pending_stock"`
	RequiresShippingMethod *bool  `json:"requires_shipping_method"`
	ServiceName            string `json:"service_name"`
	InventoryManagement    *bool  `json:"inventory_management"`
	TrackingSupport        *bool  `json:"tracking_support"`
	LocationID             int    `json:"location_id"`
}

// FulfillmentServiceAPIOp handles communication with the fulfillment service
// related methods of the Shopify API.
type FulfillmentServiceAPIOp struct {
	client *Client
}

// FulfillmentServiceListOptions fulfillment service list options
type FulfillmentServiceListOptions struct {
	Scope string `url:"scope,omitempty"`
}

// FulfillmentServiceResource fulfillment service resource
type FulfillmentServiceResource struct {
	FulfillmentService *FulfillmentService `json:"fulfillment_service"`
}

// FulfillmentServicesResource fulfillment services resource
type FulfillmentServicesResource struct {
	FulfillmentServices []FulfillmentService `json:"fulfillment_services"`
}

// List fulfillment services
func (s *FulfillmentServiceAPIOp) List(options interface{}) ([]FulfillmentService, error) {
	path := fmt.Sprintf("%s.json", fulfillmentServicesBasePath)
	resource := &FulfillmentServicesResource{}
	err := s.client.Get(path, resource, options)
	return resource.FulfillmentServices, err
}

// Get fulfillment service
func (s *FulfillmentServiceAPIOp) Get(id int, options interface{}) (*FulfillmentService, error) {
	path := fmt.Sprintf("%s/%d.json", fulfillmentServicesBasePath, id)
	resource := &FulfillmentServiceResource{}
	err := s.client.Get(path, resource, options)
	return resource.FulfillmentService, err
}

// Create a fulfillment service
func (s *FulfillmentServiceAPIOp) Create(service FulfillmentService) (*FulfillmentService, error) {
	path := fmt.Sprintf("%s.json", fulfillmentServicesBasePath)
	wrappedData := FulfillmentServiceResource{FulfillmentService: &service}
	resource := &FulfillmentServiceResource{}
	err := s.client.Post(path, wrappedData, resource)
	return resource.FulfillmentService, err
}

// Update a fulfillment service
func (s *FulfillmentServiceAPIOp) Update(service FulfillmentService) (*FulfillmentService, error) {
	path := fmt.Sprintf("%s/%d.json", fulfillmentServicesBasePath, service.ID)
	wrappedData := FulfillmentServiceResource{FulfillmentService: &service}
	resource := &FulfillmentServiceResource{}
	err := s.client.Put(path, wrappedData, resource)
	return resource.FulfillmentService, err
}

// Delete a fulfillment service
func (s *FulfillmentServiceAPIOp) Delete(id int) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", fulfillmentServicesBasePath, id))
}
