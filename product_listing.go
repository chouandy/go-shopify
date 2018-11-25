package goshopify

import (
	"fmt"
	"time"
)

const productListingsBasePath = "admin/product_listings"
const productListingsResourceName = "product_listings"

// ProductListingAPI is an interface for interfacing with the product listing
// endpoints of the Shopify API.
// See: https://help.shopify.com/en/api/reference/sales-channels/productlisting
type ProductListingAPI interface {
	List(interface{}) ([]ProductListing, error)
	ListProductIDs(interface{}) ([]int, error)
	Count(interface{}) (int, error)
	Get(int, interface{}) (*ProductListing, error)
	Publish(int) (*ProductListing, error)
	Unpublish(int) error
}

// ProductListingAPIOp handles communication with the product related methods of
// the Shopify API.
type ProductListingAPIOp struct {
	client *Client
}

// ProductListing represents a Shopify product listing
type ProductListing struct {
	ProductID   int             `json:"product_id,omitempty"`
	CreatedAt   *time.Time      `json:"created_at,omitempty"`
	UpdatedAt   *time.Time      `json:"updated_at,omitempty"`
	BodyHTML    string          `json:"body_html,omitempty"`
	Handle      string          `json:"handle,omitempty"`
	ProductType string          `json:"product_type,omitempty"`
	Title       string          `json:"title,omitempty"`
	Vendor      string          `json:"vendor,omitempty"`
	Available   *bool           `json:"available,omitempty"`
	Tags        string          `json:"tags,omitempty"`
	PublishedAt *time.Time      `json:"published_at,omitempty"`
	Variants    []Variant       `json:"variants,omitempty"`
	Images      []Image         `json:"images,omitempty"`
	Options     []ProductOption `json:"options,omitempty"`
}

// ProductListingListOptions the options for product listing list
type ProductListingListOptions struct {
	ProductIDs   []int      `url:"product_ids,omitempty,comma"`
	Limit        int        `url:"limit,omitempty"`
	Page         int        `url:"page,omitempty"`
	CollectionID int        `url:"collection_id,omitempty"`
	UpdatedAtMin *time.Time `url:"updated_at_min,omitempty"`
	Handle       string     `url:"handle,omitempty"`
}

// ProductListingResource represents the result from the product_listings/X.json endpoint
type ProductListingResource struct {
	ProductListing *ProductListing `json:"product_listing"`
}

// ProductListingsResource represents the result from the product_listings.json endpoint
type ProductListingsResource struct {
	ProductListings []ProductListing `json:"product_listings"`
}

// ProductIDsResource represents the result from the product_listings/product_ids.json endpoint
type ProductIDsResource struct {
	ProductIDs []int `json:"product_ids"`
}

// List product listings
func (s *ProductListingAPIOp) List(options interface{}) ([]ProductListing, error) {
	path := fmt.Sprintf("%s.json", productListingsBasePath)
	resource := new(ProductListingsResource)
	err := s.client.Get(path, resource, options)
	return resource.ProductListings, err
}

// ListProductIDs product ids
func (s *ProductListingAPIOp) ListProductIDs(options interface{}) ([]int, error) {
	path := fmt.Sprintf("%s/product_ids.json", productListingsBasePath)
	resource := new(ProductIDsResource)
	err := s.client.Get(path, resource, options)
	return resource.ProductIDs, err
}

// Count product listings
func (s *ProductListingAPIOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", productListingsBasePath)
	return s.client.Count(path, options)
}

// Get a product listing
func (s *ProductListingAPIOp) Get(productID int, options interface{}) (*ProductListing, error) {
	path := fmt.Sprintf("%s/%d.json", productListingsBasePath, productID)
	resource := new(ProductListingResource)
	err := s.client.Get(path, resource, options)
	return resource.ProductListing, err
}

// Publish a product
func (s *ProductListingAPIOp) Publish(productID int) (*ProductListing, error) {
	path := fmt.Sprintf("%s/%d.json", productListingsBasePath, productID)
	resource := new(ProductListingResource)
	err := s.client.Put(path, nil, resource)
	return resource.ProductListing, err
}

// Unpublish a product
func (s *ProductListingAPIOp) Unpublish(productID int) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", productListingsBasePath, productID))
}
