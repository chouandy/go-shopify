package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

// RefundAPI is an interface for interfacing with the refund endpoints
// of the Shopify API.
// https://help.shopify.com/api/reference/orders/refund
type RefundAPI interface {
	List(interface{}) ([]Refund, error)
	Get(int, interface{}) (*Refund, error)
	Calculate(Refund) (*Refund, error)
	Create(Refund) (*Refund, error)
}

// RefundsAPI is an interface for other Shopify resources
// to interface with the refund endpoints of the Shopify API.
// https://help.shopify.com/api/reference/orders/refund
type RefundsAPI interface {
	ListRefunds(int, interface{}) ([]Refund, error)
	GetRefund(int, int, interface{}) (*Refund, error)
	CalculateRefund(int, Refund) (*Refund, error)
	CreateRefund(int, Refund) (*Refund, error)
}

// RefundAPIOp handles communication with the refund
// related methods of the Shopify API.
type RefundAPIOp struct {
	client     *Client
	resource   string
	resourceID int
}

// Refund refund struct
type Refund struct {
	ID                int              `json:"id,omitempty"`
	OrderID           int              `json:"order_id,omitempty"`
	CreatedAt         *time.Time       `json:"created_at,omitempty"`
	Notify            *bool            `json:"notify,omitempty"`
	Note              string           `json:"note,omitempty"`
	UserID            *int             `json:"user_id"`
	ProcessedAt       *time.Time       `json:"processed_at,omitempty"`
	Restock           *bool            `json:"restock,omitempty"`
	AdminGraphqlAPIID string           `json:"admin_graphql_api_id,omitempty"`
	Currency          string           `json:"currency,omitempty"`
	Shipping          *Shipping        `json:"shipping,omitempty"`
	RefundLineItems   []RefundLineItem `json:"refund_line_items,omitempty"`
	Transactions      []Transaction    `json:"transactions,omitempty"`
}

// Shipping shipping
type Shipping struct {
	FullRefund        *bool            `json:"full_refund,omitempty"`
	Amount            *decimal.Decimal `json:"amount,omitempty"`
	Tax               *decimal.Decimal `json:"tax,omitempty"`
	MaximumRefundable *decimal.Decimal `json:"maximum_refundable,omitempty"`
}

// RefundLineItem refund line item
type RefundLineItem struct {
	ID                      int              `json:"id,omitempty"`
	Quantity                int              `json:"quantity,omitempty"`
	LineItemID              int              `json:"line_item_id,omitempty"`
	LocationID              *int             `json:"location_id"`
	RestockType             string           `json:"restock_type,omitempty"`
	Price                   *decimal.Decimal `json:"price,omitempty"`
	Subtotal                *decimal.Decimal `json:"subtotal,omitempty"`
	TotalTax                *decimal.Decimal `json:"total_tax,omitempty"`
	DiscountedPrice         *decimal.Decimal `json:"discounted_price,omitempty"`
	DiscountedTotalPrice    *decimal.Decimal `json:"discounted_total_price,omitempty"`
	TotalCartDiscountAmount *decimal.Decimal `json:"total_cart_discount_amount,omitempty"`
	LineItem                *LineItem        `json:"line_item,omitempty"`
}

// RefundResource represents the result from the refunds/X.json endpoint
type RefundResource struct {
	Refund *Refund `json:"refund"`
}

// RefundsResource represents the result from the refunds.json endpoint
type RefundsResource struct {
	Refunds []Refund `json:"refunds"`
}

// List refunds
func (s *RefundAPIOp) List(options interface{}) ([]Refund, error) {
	prefix := RefundPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s.json", prefix)
	resource := new(RefundsResource)
	err := s.client.Get(path, resource, options)
	return resource.Refunds, err
}

// Get individual fulfillment
func (s *RefundAPIOp) Get(fulfillmentID int, options interface{}) (*Refund, error) {
	prefix := RefundPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/%d.json", prefix, fulfillmentID)
	resource := new(RefundResource)
	err := s.client.Get(path, resource, options)
	return resource.Refund, err
}

// Calculate a new refund
func (s *RefundAPIOp) Calculate(refund Refund) (*Refund, error) {
	prefix := RefundPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/calculate.json", prefix)
	wrappedData := RefundResource{Refund: &refund}
	resource := new(RefundResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Refund, err
}

// Create a new refund
func (s *RefundAPIOp) Create(refund Refund) (*Refund, error) {
	prefix := RefundPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s.json", prefix)
	wrappedData := RefundResource{Refund: &refund}
	resource := new(RefundResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Refund, err
}
