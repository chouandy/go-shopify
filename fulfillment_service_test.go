package goshopify

import (
	"testing"

	"gopkg.in/jarcoal/httpmock.v1"
)

func fulfillmentServiceTests(t *testing.T, fulfillmentService FulfillmentService) {
	expectedInt := 755357713
	if fulfillmentService.ID != expectedInt {
		t.Errorf("FulfillmentService.ID returned %+v, expected %+v", fulfillmentService.ID, expectedInt)
	}

	expectedStr := "Shipwire App"
	if fulfillmentService.Name != expectedStr {
		t.Errorf("FulfillmentService.Name returned %+v, expected %+v", fulfillmentService.Name, expectedStr)
	}

	expectedStr = "shipwire-app"
	if fulfillmentService.Handle != expectedStr {
		t.Errorf("FulfillmentService.Handle returned %+v, expected %+v", fulfillmentService.Handle, expectedStr)
	}

	expectedBool := false
	if BoolValue(fulfillmentService.IncludePendingStock) != expectedBool {
		t.Errorf("FulfillmentService.IncludePendingStock returned %+v, expected %+v", BoolValue(fulfillmentService.IncludePendingStock), expectedStr)
	}

	expectedBool = false
	if BoolValue(fulfillmentService.RequiresShippingMethod) != expectedBool {
		t.Errorf("FulfillmentService.RequiresShippingMethod returned %+v, expected %+v", BoolValue(fulfillmentService.RequiresShippingMethod), expectedStr)
	}

	expectedStr = "Shipwire App"
	if fulfillmentService.ServiceName != expectedStr {
		t.Errorf("FulfillmentService.ServiceName returned %+v, expected %+v", fulfillmentService.ServiceName, expectedStr)
	}

	expectedBool = true
	if BoolValue(fulfillmentService.InventoryManagement) != expectedBool {
		t.Errorf("FulfillmentService.InventoryManagement returned %+v, expected %+v", BoolValue(fulfillmentService.InventoryManagement), expectedStr)
	}

	expectedBool = true
	if BoolValue(fulfillmentService.TrackingSupport) != expectedBool {
		t.Errorf("FulfillmentService.TrackingSupport returned %+v, expected %+v", BoolValue(fulfillmentService.TrackingSupport), expectedStr)
	}

	expectedInt = 48752903
	if fulfillmentService.LocationID != expectedInt {
		t.Errorf("FulfillmentService.LocationID returned %+v, expected %+v", fulfillmentService.LocationID, expectedInt)
	}
}

func TestFulfillmentServiceList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/fulfillment_services.json",
		httpmock.NewBytesResponder(200, loadFixture("fulfillment_services.json")))

	fulfillmentServices, err := client.FulfillmentService.List(nil)
	if err != nil {
		t.Errorf("FulfillmentService.List returned error: %v", err)
	}

	// Check that fulfillmentServices were parsed
	if len(fulfillmentServices) != 1 {
		t.Errorf("FulfillmentService.List got %v fulfillmentServices, expected: 1", len(fulfillmentServices))
	}

	fulfillmentServiceTests(t, fulfillmentServices[0])
}

func TestFulfillmentServiceGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/fulfillment_services/755357713.json",
		httpmock.NewBytesResponder(200, loadFixture("fulfillment_service.json")))

	fulfillmentService, err := client.FulfillmentService.Get(755357713, nil)
	if err != nil {
		t.Errorf("FulfillmentService.Get returned error: %v", err)
	}

	fulfillmentServiceTests(t, *fulfillmentService)
}

func TestFulfillmentServiceCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "https://fooshop.myshopify.com/admin/fulfillment_services.json",
		httpmock.NewBytesResponder(200, loadFixture("fulfillment_service.json")))

	fulfillmentService := FulfillmentService{
		Name:                   "Shipwire App",
		IncludePendingStock:    Bool(false),
		RequiresShippingMethod: Bool(false),
		InventoryManagement:    Bool(true),
		TrackingSupport:        Bool(true),
	}

	returnedFulfillmentService, err := client.FulfillmentService.Create(fulfillmentService)
	if err != nil {
		t.Errorf("FulfillmentService.Create returned error: %v", err)
	}

	fulfillmentServiceTests(t, *returnedFulfillmentService)
}

func TestFulfillmentServiceUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", "https://fooshop.myshopify.com/admin/fulfillment_services/755357713.json",
		httpmock.NewBytesResponder(200, loadFixture("fulfillment_service.json")))

	fulfillmentService := FulfillmentService{
		ID:                     755357713,
		Name:                   "Shipwire App",
		IncludePendingStock:    Bool(false),
		RequiresShippingMethod: Bool(false),
		InventoryManagement:    Bool(true),
		TrackingSupport:        Bool(true),
	}

	returnedFulfillmentService, err := client.FulfillmentService.Update(fulfillmentService)
	if err != nil {
		t.Errorf("FulfillmentService.Update returned error: %v", err)
	}

	fulfillmentServiceTests(t, *returnedFulfillmentService)
}

func TestFulfillmentServiceDelete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", "https://fooshop.myshopify.com/admin/fulfillment_services/755357713.json",
		httpmock.NewStringResponder(200, "{}"))

	err := client.FulfillmentService.Delete(755357713)
	if err != nil {
		t.Errorf("FulfillmentService.Delete returned error: %v", err)
	}
}
