package goshopify

import (
	"fmt"
	"time"
)

// MetafieldAPI is an interface for interfacing with the metafield endpoints
// of the Shopify API.
// https://help.shopify.com/api/reference/metafield
type MetafieldAPI interface {
	List(interface{}) ([]Metafield, error)
	Count(interface{}) (int, error)
	Get(int, interface{}) (*Metafield, error)
	Create(Metafield) (*Metafield, error)
	Update(Metafield) (*Metafield, error)
	Delete(int) error
}

// MetafieldsAPI is an interface for other Shopify resources
// to interface with the metafield endpoints of the Shopify API.
// https://help.shopify.com/api/reference/metafield
type MetafieldsAPI interface {
	ListMetafields(int, interface{}) ([]Metafield, error)
	CountMetafields(int, interface{}) (int, error)
	GetMetafield(int, int, interface{}) (*Metafield, error)
	CreateMetafield(int, Metafield) (*Metafield, error)
	UpdateMetafield(int, Metafield) (*Metafield, error)
	DeleteMetafield(int, int) error
}

// MetafieldAPIOp handles communication with the metafield
// related methods of the Shopify API.
type MetafieldAPIOp struct {
	client     *Client
	resource   string
	resourceID int
}

// Metafield represents a Shopify metafield.
type Metafield struct {
	ID            int         `json:"id,omitempty"`
	Key           string      `json:"key,omitempty"`
	Value         interface{} `json:"value,omitempty"`
	ValueType     string      `json:"value_type,omitempty"`
	Namespace     string      `json:"namespace,omitempty"`
	Description   string      `json:"description,omitempty"`
	OwnerId       int         `json:"owner_id,omitempty"`
	CreatedAt     *time.Time  `json:"created_at,omitempty"`
	UpdatedAt     *time.Time  `json:"updated_at,omitempty"`
	OwnerResource string      `json:"owner_resource,omitempty"`
}

// MetafieldResource represents the result from the metafields/X.json endpoint
type MetafieldResource struct {
	Metafield *Metafield `json:"metafield"`
}

// MetafieldsResource represents the result from the metafields.json endpoint
type MetafieldsResource struct {
	Metafields []Metafield `json:"metafields"`
}

// List metafields
func (s *MetafieldAPIOp) List(options interface{}) ([]Metafield, error) {
	prefix := MetafieldPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s.json", prefix)
	resource := new(MetafieldsResource)
	err := s.client.Get(path, resource, options)
	return resource.Metafields, err
}

// Count metafields
func (s *MetafieldAPIOp) Count(options interface{}) (int, error) {
	prefix := MetafieldPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/count.json", prefix)
	return s.client.Count(path, options)
}

// Get individual metafield
func (s *MetafieldAPIOp) Get(metafieldID int, options interface{}) (*Metafield, error) {
	prefix := MetafieldPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/%d.json", prefix, metafieldID)
	resource := new(MetafieldResource)
	err := s.client.Get(path, resource, options)
	return resource.Metafield, err
}

// Create a new metafield
func (s *MetafieldAPIOp) Create(metafield Metafield) (*Metafield, error) {
	prefix := MetafieldPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s.json", prefix)
	wrappedData := MetafieldResource{Metafield: &metafield}
	resource := new(MetafieldResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Metafield, err
}

// Update an existing metafield
func (s *MetafieldAPIOp) Update(metafield Metafield) (*Metafield, error) {
	prefix := MetafieldPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/%d.json", prefix, metafield.ID)
	wrappedData := MetafieldResource{Metafield: &metafield}
	resource := new(MetafieldResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Metafield, err
}

// Delete an existing metafield
func (s *MetafieldAPIOp) Delete(metafieldID int) error {
	prefix := MetafieldPathPrefix(s.resource, s.resourceID)
	return s.client.Delete(fmt.Sprintf("%s/%d.json", prefix, metafieldID))
}
