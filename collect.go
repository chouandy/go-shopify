package goshopify

import (
	"fmt"
	"time"
)

const collectsBasePath = "admin/collects"

// CollectAPI is an interface for interfacing with the collect endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/products/collect
type CollectAPI interface {
	List(interface{}) ([]Collect, error)
	Count(interface{}) (int, error)
}

// CollectAPIOp handles communication with the collect related methods of
// the Shopify API.
type CollectAPIOp struct {
	client *Client
}

// Collect represents a Shopify collect
type Collect struct {
	ID           int        `json:"id,omitempty"`
	CollectionID int        `json:"collection_id,omitempty"`
	ProductID    int        `json:"product_id,omitempty"`
	Featured     bool       `json:"featured,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	Position     int        `json:"position,omitempty"`
	SortValue    string     `json:"sort_value,omitempty"`
}

// CollectResource represents the result from the collects/X.json endpoint
type CollectResource struct {
	Collect *Collect `json:"collect"`
}

// CollectsResource represents the result from the collects.json endpoint
type CollectsResource struct {
	Collects []Collect `json:"collects"`
}

// List collects
func (s *CollectAPIOp) List(options interface{}) ([]Collect, error) {
	path := fmt.Sprintf("%s.json", collectsBasePath)
	resource := new(CollectsResource)
	err := s.client.Get(path, resource, options)
	return resource.Collects, err
}

// Count collects
func (s *CollectAPIOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", collectsBasePath)
	return s.client.Count(path, options)
}
