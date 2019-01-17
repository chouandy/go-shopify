package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const checkoutsBasePath = "admin/checkouts"
const checkoutsResourceName = "checkouts"

// CheckoutAPI is an interface for interfacing with the checkout endpoints
// of the Shopify API.
// See: https://help.shopify.com/en/api/reference/sales-channels/checkout
type CheckoutAPI interface {
	Create(Checkout) (*Checkout, error)
	Complete(string) (*Checkout, error)
	Get(string) (*Checkout, error)
	Update(Checkout) (*Checkout, error)
	GetShippingRates(string) ([]ShippingRate, error)
}

// CheckoutAPIOp handles communication with the checkout related methods of
// the Shopify API.
type CheckoutAPIOp struct {
	client *Client
}

// Checkout represents a Shopify checkout
type Checkout struct {
	CloneURL            string                 `json:"clone_url,omitempty"`
	CompletedAt         *time.Time             `json:"completed_at,omitempty"`
	CreatedAt           *time.Time             `json:"created_at,omitempty"`
	Currency            string                 `json:"currency,omitempty"`
	PresentmentCurrency string                 `json:"presentment_currency,omitempty"`
	CustomerID          int                    `json:"customer_id,omitempty"`
	CustomerLocale      string                 `json:"customer_locale,omitempty"`
	DeviceID            int                    `json:"device_id,omitempty"`
	DiscountCode        string                 `json:"discount_code,omitempty"`
	Email               string                 `json:"email,omitempty"`
	LegalNoticeURL      string                 `json:"legal_notice_url,omitempty"`
	LocationID          int                    `json:"location_id,omitempty"`
	Name                string                 `json:"name,omitempty"`
	Note                string                 `json:"note,omitempty"`
	NoteAttributes      map[string]interface{} `json:"note_attributes,omitempty"`
	OrderID             int                    `json:"order_id,omitempty"`
	OrderStatusURL      string                 `json:"order_status_url,omitempty"`
	Order               *Order                 `json:"order,omitempty"`
	PaymentDue          *decimal.Decimal       `json:"payment_due,omitempty"`
	PaymentURL          string                 `json:"payment_url,omitempty"`
	Phone               string                 `json:"phone,omitempty"`
	PrivacyPolicyURL    string                 `json:"privacy_policy_url,omitempty"`
	RefundPolicyURL     string                 `json:"refund_policy_url,omitempty"`
	RequiresShipping    *bool                  `json:"requires_shipping,omitempty"`
	SourceIdentifier    string                 `json:"source_identifier,omitempty"`
	SourceName          string                 `json:"source_name,omitempty"`
	SourceURL           string                 `json:"source_url,omitempty"`
	SubtotalPrice       *decimal.Decimal       `json:"subtotal_price,omitempty"`
	ShippingPolicyURL   string                 `json:"shipping_policy_url,omitempty"`
	TaxExempt           *bool                  `json:"tax_exempt,omitempty"`
	TaxesIncluded       *bool                  `json:"taxes_included,omitempty"`
	TermsOfSaleURL      string                 `json:"terms_of_sale_url,omitempty"`
	TermsOfServiceURL   string                 `json:"terms_of_service_url,omitempty"`
	Token               string                 `json:"token,omitempty"`
	TotalPrice          *decimal.Decimal       `json:"total_price,omitempty"`
	TotalTax            *decimal.Decimal       `json:"total_tax,omitempty"`
	TotalLineItemsPrice *decimal.Decimal       `json:"total_line_items_price,omitempty"`
	UpdatedAt           *time.Time             `json:"updated_at,omitempty"`
	UserID              int                    `json:"user_id,omitempty"`
	WebURL              string                 `json:"web_url,omitempty"`
	LineItems           []CheckoutLineItem     `json:"line_items"`
	GiftCards           []GiftCard             `json:"gift_cards,omitempty"`
	TaxLines            []TaxLine              `json:"tax_lines,omitempty"`
	ShippingLine        *ShippingLine          `json:"shipping_line,omitempty"`
	ShippingRate        *ShippingRate          `json:"shipping_rate,omitempty"`
	ShippingAddress     *Address               `json:"shipping_address,omitempty"`
	BillingAddress      *Address               `json:"billing_address,omitempty"`
	AppliedDiscount     *AppliedDiscount       `json:"applied_discount,omitempty"`
}

// CheckoutLineItem line item struct
type CheckoutLineItem struct {
	ID                 string                 `json:"id,omitempty"`
	Key                string                 `json:"key,omitempty"`
	ProductID          int                    `json:"product_id,omitempty"`
	VariantID          int                    `json:"variant_id,omitempty"`
	SKU                string                 `json:"sku,omitempty"`
	Vendor             string                 `json:"vendor,omitempty"`
	Title              string                 `json:"title,omitempty"`
	VariantTitle       string                 `json:"variant_title,omitempty"`
	ImageURL           string                 `json:"image_url,omitempty"`
	Taxable            *bool                  `json:"taxable,omitempty"`
	RequiresShipping   *bool                  `json:"requires_shipping,omitempty"`
	GiftCard           *bool                  `json:"gift_card,omitempty"`
	Price              *decimal.Decimal       `json:"price,omitempty"`
	CompareAtPrice     *decimal.Decimal       `json:"compare_at_price,omitempty"`
	LinePrice          *decimal.Decimal       `json:"line_price,omitempty"`
	Properties         map[string]interface{} `json:"properties,omitempty"`
	Quantity           int                    `json:"quantity,omitempty"`
	Grams              int                    `json:"grams,omitempty"`
	FulfillmentService string                 `json:"fulfillment_service,omitempty"`
	AppliedDiscounts   []AppliedDiscount      `json:"applied_discounts,omitempty"`
}

// GiftCard gift card struct
type GiftCard struct {
	ID             int              `json:"id,omitempty"`
	Code           string           `json:"code,omitempty"`
	LastCharacters string           `json:"last_characters,omitempty"`
	AmountUsed     *decimal.Decimal `json:"amount_used,omitempty"`
	Balance        *decimal.Decimal `json:"balance,omitempty"`
}

// ShippingRate shipping rate struct
type ShippingRate struct {
	ID            string           `json:"id,omitempty"`
	Price         *decimal.Decimal `json:"price,omitempty"`
	Title         string           `json:"title,omitempty"`
	Checkout      *Checkout        `json:"checkout,omitempty"`
	PhoneRequired *bool            `json:"phone_required,omitempty"`
	Handle        string           `json:"handle,omitempty"`
}

// AppliedDiscount applied discount struct
type AppliedDiscount struct {
	Amount              *decimal.Decimal `json:"amount,omitempty"`
	Title               string           `json:"title,omitempty"`
	Description         string           `json:"description,omitempty"`
	Value               string           `json:"value,omitempty"`
	ValueType           string           `json:"value_type,omitempty"`
	NonApplicableReason string           `json:"non_applicable_reason,omitempty"`
	Applicable          *bool            `json:"applicable,omitempty"`
}

// CheckoutResource represents the result from the checkouts/X.json endpoint
type CheckoutResource struct {
	Checkout *Checkout `json:"checkout"`
}

// ShippingRatesResource represents the result from the checkouts/X/shipping_rates.json endpoint
type ShippingRatesResource struct {
	ShippingRates []ShippingRate `json:"shipping_rates"`
}

// Create a new checkout
func (s *CheckoutAPIOp) Create(checkout Checkout) (*Checkout, error) {
	path := fmt.Sprintf("%s.json", checkoutsBasePath)
	wrappedData := CheckoutResource{Checkout: &checkout}
	resource := new(CheckoutResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Checkout, err
}

// Complete a checkout
func (s *CheckoutAPIOp) Complete(token string) (*Checkout, error) {
	path := fmt.Sprintf("%s/%s/complete.json", checkoutsBasePath, token)
	resource := new(CheckoutResource)
	err := s.client.Post(path, nil, resource)
	return resource.Checkout, err
}

// Get a checkout
func (s *CheckoutAPIOp) Get(token string) (*Checkout, error) {
	path := fmt.Sprintf("%s/%s.json", checkoutsBasePath, token)
	resource := new(CheckoutResource)
	err := s.client.Get(path, resource, nil)
	return resource.Checkout, err
}

// Update an existing checkout
func (s *CheckoutAPIOp) Update(checkout Checkout) (*Checkout, error) {
	path := fmt.Sprintf("%s/%s.json", checkoutsBasePath, checkout.Token)
	wrappedData := CheckoutResource{Checkout: &checkout}
	resource := new(CheckoutResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Checkout, err
}

// GetShippingRates get checkout shipping rates
func (s *CheckoutAPIOp) GetShippingRates(token string) ([]ShippingRate, error) {
	path := fmt.Sprintf("%s/%s/shipping_rates.json", checkoutsBasePath, token)
	resource := new(ShippingRatesResource)
	err := s.client.Get(path, resource, nil)
	return resource.ShippingRates, err
}
