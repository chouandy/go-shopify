package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const ordersBasePath = "admin/orders"
const ordersResourceName = "orders"

// OrderAPI is an interface for interfacing with the orders endpoints of
// the Shopify API.
// See: https://help.shopify.com/api/reference/order
type OrderAPI interface {
	List(interface{}) ([]Order, error)
	Count(interface{}) (int, error)
	Get(int, interface{}) (*Order, error)
	Create(Order) (*Order, error)
	Update(Order) (*Order, error)

	// MetafieldsAPI used for Order resource to communicate with Metafields resource
	MetafieldsAPI

	// FulfillmentsAPI used for Order resource to communicate with Fulfillments resource
	FulfillmentsAPI
}

// OrderAPIOp handles communication with the order related methods of the
// Shopify API.
type OrderAPIOp struct {
	client *Client
}

// OrderCountOptions a struct for all available order count options
type OrderCountOptions struct {
	Page              int       `url:"page,omitempty"`
	Limit             int       `url:"limit,omitempty"`
	SinceID           int       `url:"since_id,omitempty"`
	CreatedAtMin      time.Time `url:"created_at_min,omitempty"`
	CreatedAtMax      time.Time `url:"created_at_max,omitempty"`
	UpdatedAtMin      time.Time `url:"updated_at_min,omitempty"`
	UpdatedAtMax      time.Time `url:"updated_at_max,omitempty"`
	Order             string    `url:"order,omitempty"`
	Fields            string    `url:"fields,omitempty"`
	Status            string    `url:"status,omitempty"`
	FinancialStatus   string    `url:"financial_status,omitempty"`
	FulfillmentStatus string    `url:"fulfillment_status,omitempty"`
}

// OrderListOptions a struct for all available order list options.
// See: https://help.shopify.com/api/reference/order#index
type OrderListOptions struct {
	Page              int       `url:"page,omitempty"`
	Limit             int       `url:"limit,omitempty"`
	SinceID           int       `url:"since_id,omitempty"`
	Status            string    `url:"status,omitempty"`
	FinancialStatus   string    `url:"financial_status,omitempty"`
	FulfillmentStatus string    `url:"fulfillment_status,omitempty"`
	CreatedAtMin      time.Time `url:"created_at_min,omitempty"`
	CreatedAtMax      time.Time `url:"created_at_max,omitempty"`
	UpdatedAtMin      time.Time `url:"updated_at_min,omitempty"`
	UpdatedAtMax      time.Time `url:"updated_at_max,omitempty"`
	ProcessedAtMin    time.Time `url:"processed_at_min,omitempty"`
	ProcessedAtMax    time.Time `url:"processed_at_max,omitempty"`
	Fields            string    `url:"fields,omitempty"`
	Order             string    `url:"order,omitempty"`
}

// Order represents a Shopify order
type Order struct {
	ID                    int              `json:"id,omitempty"`
	Name                  string           `json:"name,omitempty"`
	Email                 string           `json:"email,omitempty"`
	CreatedAt             *time.Time       `json:"created_at,omitempty"`
	UpdatedAt             *time.Time       `json:"updated_at,omitempty"`
	CancelledAt           *time.Time       `json:"cancelled_at,omitempty"`
	ClosedAt              *time.Time       `json:"closed_at,omitempty"`
	ProcessedAt           *time.Time       `json:"processed_at,omitempty"`
	Customer              *Customer        `json:"customer,omitempty"`
	BillingAddress        *Address         `json:"billing_address,omitempty"`
	ShippingAddress       *Address         `json:"shipping_address,omitempty"`
	Currency              string           `json:"currency,omitempty"`
	TotalPrice            *decimal.Decimal `json:"total_price,omitempty"`
	SubtotalPrice         *decimal.Decimal `json:"subtotal_price,omitempty"`
	TotalDiscounts        *decimal.Decimal `json:"total_discounts,omitempty"`
	TotalLineItemsPrice   *decimal.Decimal `json:"total_line_items_price,omitempty"`
	TaxesIncluded         *bool            `json:"taxes_included,omitempty"`
	TotalTax              *decimal.Decimal `json:"total_tax,omitempty"`
	TaxLines              []TaxLine        `json:"tax_lines,omitempty"`
	TotalWeight           int              `json:"total_weight,omitempty"`
	FinancialStatus       string           `json:"financial_status,omitempty"`
	Fulfillments          []Fulfillment    `json:"fulfillments,omitempty"`
	FulfillmentStatus     string           `json:"fulfillment_status,omitempty"`
	Token                 string           `json:"token,omitempty"`
	CartToken             string           `json:"cart_token,omitempty"`
	Number                int              `json:"number,omitempty"`
	OrderNumber           int              `json:"order_number,omitempty"`
	Note                  string           `json:"note,omitempty"`
	Test                  *bool            `json:"test,omitempty"`
	BrowserIP             string           `json:"browser_ip,omitempty"`
	BuyerAcceptsMarketing *bool            `json:"buyer_accepts_marketing,omitempty"`
	CancelReason          string           `json:"cancel_reason,omitempty"`
	NoteAttributes        []NoteAttribute  `json:"note_attributes,omitempty"`
	DiscountCodes         []DiscountCode   `json:"discount_codes,omitempty"`
	LineItems             []LineItem       `json:"line_items,omitempty"`
	ShippingLines         []ShippingLine   `json:"shipping_lines,omitempty"`
	Transactions          []Transaction    `json:"transactions,omitempty"`
	AppID                 int              `json:"app_id,omitempty"`
	CustomerLocale        string           `json:"customer_locale,omitempty"`
	LandingSite           string           `json:"landing_site,omitempty"`
	ReferringSite         string           `json:"referring_site,omitempty"`
	SourceName            string           `json:"source_name,omitempty"`
	ClientDetails         *ClientDetails   `json:"client_details,omitempty"`
	Tags                  string           `json:"tags,omitempty"`
	LocationID            int              `json:"location_id,omitempty"`
	PaymentGatewayNames   []string         `json:"payment_gateway_names,omitempty"`
	ProcessingMethod      string           `json:"processing_method,omitempty"`
	Refunds               []Refund         `json:"refunds,omitempty"`
	UserID                int              `json:"user_id,omitempty"`
	OrderStatusURL        string           `json:"order_status_url,omitempty"`
	StatusURL             string           `json:"status_url,omitempty"`
	Gateway               string           `json:"gateway,omitempty"`
	Confirmed             *bool            `json:"confirmed,omitempty"`
	TotalPriceUSD         *decimal.Decimal `json:"total_price_usd,omitempty"`
	CheckoutToken         string           `json:"checkout_token,omitempty"`
	Reference             string           `json:"reference,omitempty"`
	SourceIdentifier      string           `json:"source_identifier,omitempty"`
	SourceURL             string           `json:"source_url,omitempty"`
	DeviceID              int              `json:"device_id,omitempty"`
	Phone                 string           `json:"phone,omitempty"`
	LandingSiteRef        string           `json:"landing_site_ref,omitempty"`
	CheckoutID            int              `json:"checkout_id,omitempty"`
	ContactEmail          string           `json:"contact_email,omitempty"`
	Metafields            []Metafield      `json:"metafields,omitempty"`
}

// Address address struct
type Address struct {
	ID           int     `json:"id,omitempty"`
	Address1     string  `json:"address1,omitempty"`
	Address2     string  `json:"address2,omitempty"`
	City         string  `json:"city,omitempty"`
	Company      string  `json:"company,omitempty"`
	Country      string  `json:"country,omitempty"`
	CountryCode  string  `json:"country_code,omitempty"`
	FirstName    string  `json:"first_name,omitempty"`
	LastName     string  `json:"last_name,omitempty"`
	Latitude     float64 `json:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	Name         string  `json:"name,omitempty"`
	Phone        string  `json:"phone,omitempty"`
	Province     string  `json:"province,omitempty"`
	ProvinceCode string  `json:"province_code,omitempty"`
	Zip          string  `json:"zip,omitempty"`
}

// DiscountCode discount code struct
type DiscountCode struct {
	Amount *decimal.Decimal `json:"amount,omitempty"`
	Code   string           `json:"code,omitempty"`
	Type   string           `json:"type,omitempty"`
}

// LineItem line item struct
type LineItem struct {
	ID                         int              `json:"id,omitempty"`
	ProductID                  int              `json:"product_id,omitempty"`
	VariantID                  int              `json:"variant_id,omitempty"`
	Quantity                   int              `json:"quantity,omitempty"`
	Price                      *decimal.Decimal `json:"price,omitempty"`
	TotalDiscount              *decimal.Decimal `json:"total_discount,omitempty"`
	Title                      string           `json:"title,omitempty"`
	VariantTitle               string           `json:"variant_title,omitempty"`
	Name                       string           `json:"name,omitempty"`
	SKU                        string           `json:"sku,omitempty"`
	Vendor                     string           `json:"vendor,omitempty"`
	GiftCard                   *bool            `json:"gift_card,omitempty"`
	Taxable                    *bool            `json:"taxable,omitempty"`
	FulfillmentService         string           `json:"fulfillment_service,omitempty"`
	RequiresShipping           *bool            `json:"requires_shipping,omitempty"`
	VariantInventoryManagement string           `json:"variant_inventory_management,omitempty"`
	PreTaxPrice                *decimal.Decimal `json:"pre_tax_price,omitempty"`
	Properties                 []NoteAttribute  `json:"properties,omitempty"`
	ProductExists              *bool            `json:"product_exists,omitempty"`
	FulfillableQuantity        int              `json:"fulfillable_quantity,omitempty"`
	Grams                      int              `json:"grams,omitempty"`
	FulfillmentStatus          string           `json:"fulfillment_status,omitempty"`
	TaxLines                   []TaxLine        `json:"tax_lines,omitempty"`
	OriginLocation             *Address         `json:"origin_location,omitempty"`
	DestinationLocation        *Address         `json:"destination_location,omitempty"`
}

// LineItemProperty line item property struct
type LineItemProperty struct {
	Message string `json:"message"`
}

// NoteAttribute note attribute struct
type NoteAttribute struct {
	Name  string      `json:"name,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// OrderResource represents the result from the orders/X.json endpoint
type OrderResource struct {
	Order *Order `json:"order"`
}

// OrdersResource represents the result from the orders.json endpoint
type OrdersResource struct {
	Orders []Order `json:"orders"`
}

// PaymentDetails payment details struct
type PaymentDetails struct {
	AVSResultCode     string `json:"avs_result_code,omitempty"`
	CreditCardBin     string `json:"credit_card_bin,omitempty"`
	CVVResultCode     string `json:"cvv_result_code,omitempty"`
	CreditCardNumber  string `json:"credit_card_number,omitempty"`
	CreditCardCompany string `json:"credit_card_company,omitempty"`
}

// ShippingLine shipping line struct
type ShippingLine struct {
	ID                            int              `json:"id,omitempty"`
	Title                         string           `json:"title,omitempty"`
	Price                         *decimal.Decimal `json:"price,omitempty"`
	Code                          string           `json:"code,omitempty"`
	Source                        string           `json:"source,omitempty"`
	Phone                         string           `json:"phone,omitempty"`
	RequestedFulfillmentServiceID string           `json:"requested_fulfillment_service_id,omitempty"`
	DeliveryCategory              string           `json:"delivery_category,omitempty"`
	CarrierIdentifier             string           `json:"carrier_identifier,omitempty"`
	TaxLines                      []TaxLine        `json:"tax_lines,omitempty"`
	Handle                        string           `json:"handle,omitempty"`
}

// TaxLine tax line struct
type TaxLine struct {
	Title string           `json:"title,omitempty"`
	Price *decimal.Decimal `json:"price,omitempty"`
	Rate  *decimal.Decimal `json:"rate,omitempty"`
}

// Transaction transaction struct
type Transaction struct {
	ID             int              `json:"id,omitempty"`
	OrderID        int              `json:"order_id,omitempty"`
	Amount         *decimal.Decimal `json:"amount,omitempty"`
	Kind           string           `json:"kind,omitempty"`
	Gateway        string           `json:"gateway,omitempty"`
	Status         string           `json:"status,omitempty"`
	Message        string           `json:"message,omitempty"`
	CreatedAt      *time.Time       `json:"created_at,omitempty"`
	Test           *bool            `json:"test,omitempty"`
	Authorization  string           `json:"authorization,omitempty"`
	Currency       string           `json:"currency,omitempty"`
	LocationID     *int             `json:"location_id,omitempty"`
	UserID         *int             `json:"user_id,omitempty"`
	ParentID       *int             `json:"parent_id,omitempty"`
	DeviceID       *int             `json:"device_id,omitempty"`
	ErrorCode      string           `json:"error_code,omitempty"`
	SourceName     string           `json:"source_name,omitempty"`
	PaymentDetails *PaymentDetails  `json:"payment_details,omitempty"`
}

// ClientDetails client details struct
type ClientDetails struct {
	AcceptLanguage string `json:"accept_language,omitempty"`
	BrowserHeight  int    `json:"browser_height,omitempty"`
	BrowserIP      string `json:"browser_ip,omitempty"`
	BrowserWidth   int    `json:"browser_width,omitempty"`
	SessionHash    string `json:"session_hash,omitempty"`
	UserAgent      string `json:"user_agent,omitempty"`
}

// Refund refund struct
type Refund struct {
	ID              int              `json:"id,omitempty"`
	OrderID         int              `json:"order_id,omitempty"`
	CreatedAt       *time.Time       `json:"created_at,omitempty"`
	Note            string           `json:"note,omitempty"`
	Restock         *bool            `json:"restock,omitempty"`
	UserID          int              `json:"user_id,omitempty"`
	RefundLineItems []RefundLineItem `json:"refund_line_items,omitempty"`
	Transactions    []Transaction    `json:"transactions,omitempty"`
}

// RefundLineItem refund line item
type RefundLineItem struct {
	ID         int              `json:"id,omitempty"`
	Quantity   int              `json:"quantity,omitempty"`
	LineItemID int              `json:"line_item_id,omitempty"`
	LineItem   *LineItem        `json:"line_item,omitempty"`
	Subtotal   *decimal.Decimal `json:"subtotal,omitempty"`
	TotalTax   *decimal.Decimal `json:"total_tax,omitempty"`
}

// List orders
func (s *OrderAPIOp) List(options interface{}) ([]Order, error) {
	path := fmt.Sprintf("%s.json", ordersBasePath)
	resource := new(OrdersResource)
	err := s.client.Get(path, resource, options)
	return resource.Orders, err
}

// Count orders
func (s *OrderAPIOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", ordersBasePath)
	return s.client.Count(path, options)
}

// Get individual order
func (s *OrderAPIOp) Get(orderID int, options interface{}) (*Order, error) {
	path := fmt.Sprintf("%s/%d.json", ordersBasePath, orderID)
	resource := new(OrderResource)
	err := s.client.Get(path, resource, options)
	return resource.Order, err
}

// Create order
func (s *OrderAPIOp) Create(order Order) (*Order, error) {
	path := fmt.Sprintf("%s.json", ordersBasePath)
	wrappedData := OrderResource{Order: &order}
	resource := new(OrderResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Order, err
}

// Update order
func (s *OrderAPIOp) Update(order Order) (*Order, error) {
	path := fmt.Sprintf("%s/%d.json", ordersBasePath, order.ID)
	wrappedData := OrderResource{Order: &order}
	resource := new(OrderResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Order, err
}

// ListMetafields list metafields for an order
func (s *OrderAPIOp) ListMetafields(orderID int, options interface{}) ([]Metafield, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return metafieldAPI.List(options)
}

// CountMetafields count metafields for an order
func (s *OrderAPIOp) CountMetafields(orderID int, options interface{}) (int, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return metafieldAPI.Count(options)
}

// GetMetafield get individual metafield for an order
func (s *OrderAPIOp) GetMetafield(orderID int, metafieldID int, options interface{}) (*Metafield, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return metafieldAPI.Get(metafieldID, options)
}

// CreateMetafield create a new metafield for an order
func (s *OrderAPIOp) CreateMetafield(orderID int, metafield Metafield) (*Metafield, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return metafieldAPI.Create(metafield)
}

// UpdateMetafield update an existing metafield for an order
func (s *OrderAPIOp) UpdateMetafield(orderID int, metafield Metafield) (*Metafield, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return metafieldAPI.Update(metafield)
}

// DeleteMetafield delete an existing metafield for an order
func (s *OrderAPIOp) DeleteMetafield(orderID int, metafieldID int) error {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return metafieldAPI.Delete(metafieldID)
}

// ListFulfillments list fulfillments for an order
func (s *OrderAPIOp) ListFulfillments(orderID int, options interface{}) ([]Fulfillment, error) {
	fulfillmentAPI := &FulfillmentAPIOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return fulfillmentAPI.List(options)
}

// CountFulfillments count fulfillments for an order
func (s *OrderAPIOp) CountFulfillments(orderID int, options interface{}) (int, error) {
	fulfillmentAPI := &FulfillmentAPIOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return fulfillmentAPI.Count(options)
}

// GetFulfillment get individual fulfillment for an order
func (s *OrderAPIOp) GetFulfillment(orderID int, fulfillmentID int, options interface{}) (*Fulfillment, error) {
	fulfillmentAPI := &FulfillmentAPIOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return fulfillmentAPI.Get(fulfillmentID, options)
}

// CreateFulfillment create a new fulfillment for an order
func (s *OrderAPIOp) CreateFulfillment(orderID int, fulfillment Fulfillment) (*Fulfillment, error) {
	fulfillmentAPI := &FulfillmentAPIOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return fulfillmentAPI.Create(fulfillment)
}

// UpdateFulfillment update an existing fulfillment for an order
func (s *OrderAPIOp) UpdateFulfillment(orderID int, fulfillment Fulfillment) (*Fulfillment, error) {
	fulfillmentAPI := &FulfillmentAPIOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return fulfillmentAPI.Update(fulfillment)
}

// CompleteFulfillment complete an existing fulfillment for an order
func (s *OrderAPIOp) CompleteFulfillment(orderID int, fulfillmentID int) (*Fulfillment, error) {
	fulfillmentAPI := &FulfillmentAPIOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return fulfillmentAPI.Complete(fulfillmentID)
}

// OpenFulfillment open an existing fulfillment for an order
func (s *OrderAPIOp) OpenFulfillment(orderID int, fulfillmentID int) (*Fulfillment, error) {
	fulfillmentAPI := &FulfillmentAPIOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return fulfillmentAPI.Open(fulfillmentID)
}

// CancelFulfillment cancel an existing fulfillment for an order
func (s *OrderAPIOp) CancelFulfillment(orderID int, fulfillmentID int) (*Fulfillment, error) {
	fulfillmentAPI := &FulfillmentAPIOp{client: s.client, resource: ordersResourceName, resourceID: orderID}
	return fulfillmentAPI.Cancel(fulfillmentID)
}
