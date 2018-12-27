package goshopify

import (
	"reflect"
	"testing"

	"github.com/shopspring/decimal"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func RefundTests(t *testing.T, refund Refund) {
	// Check that ID is assigned to the returned refund
	expectedInt := 929361462
	if refund.ID != expectedInt {
		t.Errorf("Refund.ID returned %+v, expected %+v", refund.ID, expectedInt)
	}
}

func TestRefundList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/orders/1/refunds.json",
		httpmock.NewStringResponder(200, `{"refunds": [{"id":1},{"id":2}]}`))

	refundAPI := &RefundAPIOp{client: client, resource: ordersResourceName, resourceID: 1}

	refunds, err := refundAPI.List(nil)
	if err != nil {
		t.Errorf("Refund.List returned error: %v", err)
	}

	expected := []Refund{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(refunds, expected) {
		t.Errorf("Refund.List returned %+v, expected %+v", refunds, expected)
	}
}

func TestRefundGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/orders/1/refunds/1.json",
		httpmock.NewStringResponder(200, `{"refund": {"id":1}}`))

	refundAPI := &RefundAPIOp{client: client, resource: ordersResourceName, resourceID: 1}

	refund, err := refundAPI.Get(1, nil)
	if err != nil {
		t.Errorf("Refund.Get returned error: %v", err)
	}

	expected := &Refund{ID: 1}
	if !reflect.DeepEqual(refund, expected) {
		t.Errorf("Refund.Get returned %+v, expected %+v", refund, expected)
	}
}

func TestRefundCalculate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "https://fooshop.myshopify.com/admin/orders/1/refunds/calculate.json",
		httpmock.NewBytesResponder(200, loadFixture("refund_calculate.json")))

	refundAPI := &RefundAPIOp{client: client, resource: ordersResourceName, resourceID: 1}

	refund := Refund{
		Currency: "USD",
		Notify:   Bool(true),
		Note:     "wrong size",
		Shipping: &Shipping{
			FullRefund: Bool(true),
		},
		RefundLineItems: []RefundLineItem{
			{
				LineItemID:  518995019,
				Quantity:    1,
				RestockType: "no_restock",
			},
		},
	}

	returnedRefund, err := refundAPI.Calculate(refund)
	if err != nil {
		t.Errorf("Refund.Calculate returned error: %v", err)
	}

	expectedInt := 518995019
	if returnedRefund.RefundLineItems[0].LineItemID != expectedInt {
		t.Errorf("Refund.RefundLineItems[0].LineItemID returned %+v, expected %+v", refund.ID, expectedInt)
	}
	expectedInt = 1
	if returnedRefund.RefundLineItems[0].Quantity != expectedInt {
		t.Errorf("Refund.RefundLineItems[0].Quantity returned %+v, expected %+v", refund.ID, expectedInt)
	}
	expectedStr := "no_restock"
	if returnedRefund.RefundLineItems[0].RestockType != expectedStr {
		t.Errorf("Refund.RefundLineItems[0].RestockType returned %+v, expected %+v", refund.ID, expectedInt)
	}
}

func TestRefundCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "https://fooshop.myshopify.com/admin/orders/1/refunds.json",
		httpmock.NewBytesResponder(200, loadFixture("refund.json")))

	refundAPI := &RefundAPIOp{client: client, resource: ordersResourceName, resourceID: 1}

	amount := decimal.NewFromFloat(41.94)
	refund := Refund{
		Currency: "USD",
		Notify:   Bool(true),
		Note:     "wrong size",
		Shipping: &Shipping{
			FullRefund: Bool(true),
		},
		RefundLineItems: []RefundLineItem{
			{
				LineItemID:  518995019,
				Quantity:    1,
				RestockType: "return",
				LocationID:  Int(487838322),
			},
		},
		Transactions: []Transaction{
			{
				ParentID: Int(801038806),
				Amount:   &amount,
				Kind:     "refund",
				Gateway:  "bogus",
			},
		},
	}

	returnedRefund, err := refundAPI.Create(refund)
	if err != nil {
		t.Errorf("Refund.Create returned error: %v", err)
	}

	RefundTests(t, *returnedRefund)
}
