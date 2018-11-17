package goshopify

import (
	"fmt"
	"time"
)

const productsBasePath = "admin/products"
const productsResourceName = "products"

// ProductService is an interface for interfacing with the product endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/product
type ProductService interface {
	List(interface{}) ([]Product, error)
	Count(interface{}) (int, error)
	Get(int, interface{}) (*Product, error)
	Create(Product) (*Product, error)
	Update(Product) (*Product, error)
	Delete(int) error

	// MetafieldsAPI used for Product resource to communicate with Metafields resource
	MetafieldsAPI
}

// ProductServiceOp handles communication with the product related methods of
// the Shopify API.
type ProductServiceOp struct {
	client *Client
}

// Product represents a Shopify product
type Product struct {
	ID                             int             `json:"id,omitempty"`
	Title                          string          `json:"title,omitempty"`
	BodyHTML                       string          `json:"body_html,omitempty"`
	Vendor                         string          `json:"vendor,omitempty"`
	ProductType                    string          `json:"product_type,omitempty"`
	Handle                         string          `json:"handle,omitempty"`
	CreatedAt                      *time.Time      `json:"created_at,omitempty"`
	UpdatedAt                      *time.Time      `json:"updated_at,omitempty"`
	PublishedAt                    *time.Time      `json:"published_at,omitempty"`
	PublishedScope                 string          `json:"published_scope,omitempty"`
	Tags                           string          `json:"tags,omitempty"`
	Options                        []ProductOption `json:"options,omitempty"`
	Variants                       []Variant       `json:"variants,omitempty"`
	Image                          *Image          `json:"image,omitempty"`
	Images                         []Image         `json:"images,omitempty"`
	TemplateSuffix                 string          `json:"template_suffix,omitempty"`
	MetafieldsGlobalTitleTag       string          `json:"metafields_global_title_tag,omitempty"`
	MetafieldsGlobalDescriptionTag string          `json:"metafields_global_description_tag,omitempty"`
	Metafields                     []Metafield     `json:"metafields,omitempty"`
	AdminGraphqlAPIID              string          `json:"admin_graphql_api_id,omitempty"`
}

// ProductOption the options provided by Shopify
type ProductOption struct {
	ID        int      `json:"id,omitempty"`
	ProductID int      `json:"product_id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Position  int      `json:"position,omitempty"`
	Values    []string `json:"values,omitempty"`
}

// ProductResource represents the result from the products/X.json endpoint
type ProductResource struct {
	Product *Product `json:"product"`
}

// ProductsResource represents the result from the products.json endpoint
type ProductsResource struct {
	Products []Product `json:"products"`
}

// List products
func (s *ProductServiceOp) List(options interface{}) ([]Product, error) {
	path := fmt.Sprintf("%s.json", productsBasePath)
	resource := new(ProductsResource)
	err := s.client.Get(path, resource, options)
	return resource.Products, err
}

// Count products
func (s *ProductServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", productsBasePath)
	return s.client.Count(path, options)
}

// Get individual product
func (s *ProductServiceOp) Get(productID int, options interface{}) (*Product, error) {
	path := fmt.Sprintf("%s/%d.json", productsBasePath, productID)
	resource := new(ProductResource)
	err := s.client.Get(path, resource, options)
	return resource.Product, err
}

// Create a new product
func (s *ProductServiceOp) Create(product Product) (*Product, error) {
	path := fmt.Sprintf("%s.json", productsBasePath)
	wrappedData := ProductResource{Product: &product}
	resource := new(ProductResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Product, err
}

// Update an existing product
func (s *ProductServiceOp) Update(product Product) (*Product, error) {
	path := fmt.Sprintf("%s/%d.json", productsBasePath, product.ID)
	wrappedData := ProductResource{Product: &product}
	resource := new(ProductResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Product, err
}

// Delete an existing product
func (s *ProductServiceOp) Delete(productID int) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", productsBasePath, productID))
}

// ListMetafields list metafields for a product
func (s *ProductServiceOp) ListMetafields(productID int, options interface{}) ([]Metafield, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: productsResourceName, resourceID: productID}
	return metafieldAPI.List(options)
}

// CountMetafields count metafields for a product
func (s *ProductServiceOp) CountMetafields(productID int, options interface{}) (int, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: productsResourceName, resourceID: productID}
	return metafieldAPI.Count(options)
}

// GetMetafield get individual metafield for a product
func (s *ProductServiceOp) GetMetafield(productID int, metafieldID int, options interface{}) (*Metafield, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: productsResourceName, resourceID: productID}
	return metafieldAPI.Get(metafieldID, options)
}

// CreateMetafield create a new metafield for a product
func (s *ProductServiceOp) CreateMetafield(productID int, metafield Metafield) (*Metafield, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: productsResourceName, resourceID: productID}
	return metafieldAPI.Create(metafield)
}

// UpdateMetafield update an existing metafield for a product
func (s *ProductServiceOp) UpdateMetafield(productID int, metafield Metafield) (*Metafield, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: productsResourceName, resourceID: productID}
	return metafieldAPI.Update(metafield)
}

// DeleteMetafield delete an existing metafield for a product
func (s *ProductServiceOp) DeleteMetafield(productID int, metafieldID int) error {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: productsResourceName, resourceID: productID}
	return metafieldAPI.Delete(metafieldID)
}
