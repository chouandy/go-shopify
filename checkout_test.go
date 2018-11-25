package goshopify

import (
	"reflect"
	"testing"

	"gopkg.in/jarcoal/httpmock.v1"
)

func checkoutTests(t *testing.T, checkout Checkout) {
	// Check that Token is assigned to the returned checkout
	expectedStr := "aa8296288254541a747d50794cee3249"
	if checkout.Token != expectedStr {
		t.Errorf("Checkout.Token returned %+v, expected %+v", checkout.Token, expectedStr)
	}
}

func shippingRateTests(t *testing.T, shippingRate ShippingRate) {
	expectedStr := "shopify-Standard%20Shipping-10.99"
	if shippingRate.ID != expectedStr {
		t.Errorf("Checkout.Token returned %+v, expected %+v", shippingRate.ID, expectedStr)
	}

	expectedStr = "Standard Shipping"
	if shippingRate.Title != expectedStr {
		t.Errorf("Checkout.Token returned %+v, expected %+v", shippingRate.Title, expectedStr)
	}

	expectedBool := false
	if BoolValue(shippingRate.PhoneRequired) != expectedBool {
		t.Errorf("Checkout.Token returned %+v, expected %+v", shippingRate.PhoneRequired, expectedStr)
	}

	expectedStr = "shopify-Standard%20Shipping-10.99"
	if shippingRate.Handle != expectedStr {
		t.Errorf("Checkout.Token returned %+v, expected %+v", shippingRate.Handle, expectedStr)
	}
}

func TestCheckoutCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "https://fooshop.myshopify.com/admin/checkouts.json",
		httpmock.NewBytesResponder(200, loadFixture("checkout.json")))

	checkout := Checkout{
		LineItems: []CheckoutLineItem{
			{
				VariantID: 13142253437041,
				Quantity:  1,
			},
		},
	}

	returnedCheckout, err := client.Checkout.Create(checkout)
	if err != nil {
		t.Errorf("Checkout.Create returned error: %v", err)
	}

	checkoutTests(t, *returnedCheckout)
}

func TestCheckoutComplete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "https://fooshop.myshopify.com/admin/checkouts/aa8296288254541a747d50794cee3249/complete.json",
		httpmock.NewBytesResponder(200, loadFixture("checkout.json")))

	checkout, err := client.Checkout.Complete("aa8296288254541a747d50794cee3249")
	if err != nil {
		t.Errorf("Checkout.Complete returned error: %v", err)
	}

	checkoutTests(t, *checkout)
}

func TestCheckoutGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/checkouts/aa8296288254541a747d50794cee3249.json",
		httpmock.NewStringResponder(200, `{"checkout": {"token":"aa8296288254541a747d50794cee3249"}}`))

	checkout, err := client.Checkout.Get("aa8296288254541a747d50794cee3249")
	if err != nil {
		t.Errorf("Checkout.Get returned error: %v", err)
	}

	expected := &Checkout{Token: "aa8296288254541a747d50794cee3249"}
	if !reflect.DeepEqual(checkout, expected) {
		t.Errorf("Checkout.Get returned %+v, expected %+v", checkout, expected)
	}
}

func TestCheckoutUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", "https://fooshop.myshopify.com/admin/checkouts/aa8296288254541a747d50794cee3249.json",
		httpmock.NewBytesResponder(200, loadFixture("checkout.json")))

	checkout := Checkout{
		Token: "aa8296288254541a747d50794cee3249",
	}

	returnedCheckout, err := client.Checkout.Update(checkout)
	if err != nil {
		t.Errorf("Checkout.Update returned error: %v", err)
	}

	checkoutTests(t, *returnedCheckout)
}

func TestCheckoutGetShippingRates(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/checkouts/aa8296288254541a747d50794cee3249/shipping_rates.json",
		httpmock.NewBytesResponder(200, loadFixture("shipping_rates.json")))

	shippingRates, err := client.Checkout.GetShippingRates("aa8296288254541a747d50794cee3249")
	if err != nil {
		t.Errorf("Checkout.Get returned error: %v", err)
	}

	for _, shippingRate := range shippingRates {
		shippingRateTests(t, shippingRate)
	}
}
