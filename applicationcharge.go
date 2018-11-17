package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const applicationChargesBasePath = "admin/application_charges"

// ApplicationChargeAPI is an interface for interacting with the
// ApplicationCharge endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/billing/applicationcharge
type ApplicationChargeAPI interface {
	Create(ApplicationCharge) (*ApplicationCharge, error)
	Get(int, interface{}) (*ApplicationCharge, error)
	List(interface{}) ([]ApplicationCharge, error)
	Activate(ApplicationCharge) (*ApplicationCharge, error)
}

// ApplicationChargeAPIOp application charge service op
type ApplicationChargeAPIOp struct {
	client *Client
}

// ApplicationCharge application charge
type ApplicationCharge struct {
	ID                 int              `json:"id"`
	Name               string           `json:"name"`
	APIClientID        int              `json:"api_client_id"`
	Price              *decimal.Decimal `json:"price"`
	Status             string           `json:"status"`
	ReturnURL          string           `json:"return_url"`
	Test               *bool            `json:"test"`
	CreatedAt          *time.Time       `json:"created_at"`
	UpdatedAt          *time.Time       `json:"updated_at"`
	ChargeType         *string          `json:"charge_type"`
	DecoratedReturnURL string           `json:"decorated_return_url"`
	ConfirmationURL    string           `json:"confirmation_url"`
}

// ApplicationChargeResource represents the result from the
// admin/application_charges{/X{/activate.json}.json}.json endpoints.
type ApplicationChargeResource struct {
	Charge *ApplicationCharge `json:"application_charge"`
}

// ApplicationChargesResource represents the result from the
// admin/application_charges.json endpoint.
type ApplicationChargesResource struct {
	Charges []ApplicationCharge `json:"application_charges"`
}

// Create creates new application charge.
func (a *ApplicationChargeAPIOp) Create(charge ApplicationCharge) (*ApplicationCharge, error) {
	path := fmt.Sprintf("%s.json", applicationChargesBasePath)
	resource := &ApplicationChargeResource{}
	return resource.Charge, a.client.Post(path, ApplicationChargeResource{Charge: &charge}, resource)
}

// Get gets individual application charge.
func (a *ApplicationChargeAPIOp) Get(chargeID int, options interface{}) (*ApplicationCharge, error) {
	path := fmt.Sprintf("%s/%d.json", applicationChargesBasePath, chargeID)
	resource := &ApplicationChargeResource{}
	return resource.Charge, a.client.Get(path, resource, options)
}

// List gets all application charges.
func (a *ApplicationChargeAPIOp) List(options interface{}) ([]ApplicationCharge, error) {
	path := fmt.Sprintf("%s.json", applicationChargesBasePath)
	resource := &ApplicationChargesResource{}
	return resource.Charges, a.client.Get(path, resource, options)
}

// Activate activates application charge.
func (a *ApplicationChargeAPIOp) Activate(charge ApplicationCharge) (*ApplicationCharge, error) {
	path := fmt.Sprintf("%s/%d/activate.json", applicationChargesBasePath, charge.ID)
	resource := &ApplicationChargeResource{}
	return resource.Charge, a.client.Post(path, ApplicationChargeResource{Charge: &charge}, resource)
}
