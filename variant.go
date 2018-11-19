package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const variantsBasePath = "admin/variants"

// VariantAPI is an interface for interacting with the variant endpoints
// of the Shopify API.
// See https://help.shopify.com/api/reference/product_variant
type VariantAPI interface {
	List(int, interface{}) ([]Variant, error)
	Count(int, interface{}) (int, error)
	Get(int, interface{}) (*Variant, error)
	Create(int, Variant) (*Variant, error)
	Update(Variant) (*Variant, error)
	Delete(int, int) error
}

// VariantAPIOp handles communication with the variant related methods of
// the Shopify API.
type VariantAPIOp struct {
	client *Client
}

// Variant represents a Shopify variant
type Variant struct {
	ID                   int              `json:"id,omitempty"`
	ProductID            int              `json:"product_id,omitempty"`
	Title                string           `json:"title,omitempty"`
	Sku                  string           `json:"sku,omitempty"`
	Position             int              `json:"position,omitempty"`
	Grams                int              `json:"grams,omitempty"`
	InventoryPolicy      string           `json:"inventory_policy,omitempty"`
	Price                *decimal.Decimal `json:"price,omitempty"`
	CompareAtPrice       *decimal.Decimal `json:"compare_at_price,omitempty"`
	FulfillmentService   string           `json:"fulfillment_service,omitempty"`
	InventoryManagement  string           `json:"inventory_management,omitempty"`
	InventoryItemId      int              `json:"inventory_item_id,omitempty"`
	Option1              string           `json:"option1,omitempty"`
	Option2              string           `json:"option2,omitempty"`
	Option3              string           `json:"option3,omitempty"`
	CreatedAt            *time.Time       `json:"created_at,omitempty"`
	UpdatedAt            *time.Time       `json:"updated_at,omitempty"`
	Taxable              bool             `json:"taxable,omitempty"`
	Barcode              string           `json:"barcode,omitempty"`
	ImageID              int              `json:"image_id,omitempty"`
	InventoryQuantity    int              `json:"inventory_quantity,omitempty"`
	Weight               *decimal.Decimal `json:"weight,omitempty"`
	WeightUnit           string           `json:"weight_unit,omitempty"`
	OldInventoryQuantity int              `json:"old_inventory_quantity,omitempty"`
	RequireShipping      bool             `json:"requires_shipping,omitempty"`
	AdminGraphqlAPIID    string           `json:"admin_graphql_api_id,omitempty"`
}

// VariantResource represents the result from the variants/X.json endpoint
type VariantResource struct {
	Variant *Variant `json:"variant"`
}

// VariantsResource represents the result from the products/X/variants.json endpoint
type VariantsResource struct {
	Variants []Variant `json:"variants"`
}

// List variants
func (s *VariantAPIOp) List(productID int, options interface{}) ([]Variant, error) {
	path := fmt.Sprintf("%s/%d/variants.json", productsBasePath, productID)
	resource := new(VariantsResource)
	err := s.client.Get(path, resource, options)
	return resource.Variants, err
}

// Count variants
func (s *VariantAPIOp) Count(productID int, options interface{}) (int, error) {
	path := fmt.Sprintf("%s/%d/variants/count.json", productsBasePath, productID)
	return s.client.Count(path, options)
}

// Get individual variant
func (s *VariantAPIOp) Get(variantID int, options interface{}) (*Variant, error) {
	path := fmt.Sprintf("%s/%d.json", variantsBasePath, variantID)
	resource := new(VariantResource)
	err := s.client.Get(path, resource, options)
	return resource.Variant, err
}

// Create a new variant
func (s *VariantAPIOp) Create(productID int, variant Variant) (*Variant, error) {
	path := fmt.Sprintf("%s/%d/variants.json", productsBasePath, productID)
	wrappedData := VariantResource{Variant: &variant}
	resource := new(VariantResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Variant, err
}

// Update existing variant
func (s *VariantAPIOp) Update(variant Variant) (*Variant, error) {
	path := fmt.Sprintf("%s/%d.json", variantsBasePath, variant.ID)
	wrappedData := VariantResource{Variant: &variant}
	resource := new(VariantResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Variant, err
}

// Delete an existing product
func (s *VariantAPIOp) Delete(productID int, variantID int) error {
	return s.client.Delete(fmt.Sprintf("%s/%d/variants/%d.json", productsBasePath, productID, variantID))
}
