package goshopify

import (
	"reflect"
	"testing"
)

func TestShopFullName(t *testing.T) {
	cases := []struct {
		in, expected string
	}{
		{"myshop", "myshop.myshopify.com"},
		{"myshop.", "myshop.myshopify.com"},
		{" myshop", "myshop.myshopify.com"},
		{"myshop ", "myshop.myshopify.com"},
		{"myshop \n", "myshop.myshopify.com"},
		{"myshop.myshopify.com", "myshop.myshopify.com"},
	}

	for _, c := range cases {
		actual := ShopFullName(c.in)
		if actual != c.expected {
			t.Errorf("ShopFullName(%s): expected %s, actual %s", c.in, c.expected, actual)
		}
	}
}

func TestShopShortName(t *testing.T) {
	cases := []struct {
		in, expected string
	}{
		{"myshop", "myshop"},
		{"myshop.", "myshop"},
		{" myshop", "myshop"},
		{"myshop ", "myshop"},
		{"myshop \n", "myshop"},
		{"myshop.myshopify.com", "myshop"},
		{".myshop.myshopify.com.", "myshop"},
	}

	for _, c := range cases {
		actual := ShopShortName(c.in)
		if actual != c.expected {
			t.Errorf("ShopShortName(%s): expected %s, actual %s", c.in, c.expected, actual)
		}
	}
}

func TestShopBaseURL(t *testing.T) {
	cases := []struct {
		in, expected string
	}{
		{"myshop", "https://myshop.myshopify.com"},
		{"myshop.", "https://myshop.myshopify.com"},
		{" myshop", "https://myshop.myshopify.com"},
		{"myshop ", "https://myshop.myshopify.com"},
		{"myshop \n", "https://myshop.myshopify.com"},
		{"myshop.myshopify.com", "https://myshop.myshopify.com"},
	}

	for _, c := range cases {
		actual := ShopBaseURL(c.in)
		if actual != c.expected {
			t.Errorf("ShopBaseURL(%s): expected %s, actual %s", c.in, c.expected, actual)
		}
	}
}

func TestMetafieldPathPrefix(t *testing.T) {
	cases := []struct {
		resource   string
		resourceID int
		expected   string
	}{
		{"", 0, "admin/metafields"},
		{"products", 123, "admin/products/123/metafields"},
	}

	for _, c := range cases {
		actual := MetafieldPathPrefix(c.resource, c.resourceID)
		if actual != c.expected {
			t.Errorf("MetafieldPathPrefix(%s, %d): expected %s, actual %s", c.resource, c.resourceID, c.expected, actual)
		}
	}
}

func TestFulfillmentPathPrefix(t *testing.T) {
	cases := []struct {
		resource   string
		resourceID int
		expected   string
	}{
		{"", 0, "admin/fulfillments"},
		{"orders", 123, "admin/orders/123/fulfillments"},
	}

	for _, c := range cases {
		actual := FulfillmentPathPrefix(c.resource, c.resourceID)
		if actual != c.expected {
			t.Errorf("FulfillmentPathPrefix(%s, %d): expected %s, actual %s", c.resource, c.resourceID, c.expected, actual)
		}
	}
}

func TestRefundPathPrefix(t *testing.T) {
	cases := []struct {
		resource   string
		resourceID int
		expected   string
	}{
		{"", 0, "admin/refunds"},
		{"orders", 123, "admin/orders/123/refunds"},
	}

	for _, c := range cases {
		actual := RefundPathPrefix(c.resource, c.resourceID)
		if actual != c.expected {
			t.Errorf("RefundPathPrefix(%s, %d): expected %s, actual %s", c.resource, c.resourceID, c.expected, actual)
		}
	}
}

func TestBool(t *testing.T) {
	cases := []struct {
		value    bool
		expected string
	}{
		{
			true,
			"*bool",
		},
		{
			false,
			"*bool",
		},
	}

	for _, c := range cases {
		actual := reflect.TypeOf(Bool(c.value)).String()
		if actual != c.expected {
			t.Errorf("Bool(%v): expected %s, actual %s", c.value, c.expected, actual)
		}
	}
}

func TestBoolValue(t *testing.T) {
	cases := []struct {
		value    *bool
		expected string
	}{
		{
			Bool(true),
			"bool",
		},
		{
			Bool(false),
			"bool",
		},
		{
			nil,
			"bool",
		},
	}

	for _, c := range cases {
		actual := reflect.TypeOf(BoolValue(c.value)).String()
		if actual != c.expected {
			t.Errorf("BoolValue(%v): expected %s, actual %s", c.value, c.expected, actual)
		}
	}
}

func TestString(t *testing.T) {
	cases := []struct {
		value    string
		expected string
	}{
		{
			"string",
			"*string",
		},
	}

	for _, c := range cases {
		actual := reflect.TypeOf(String(c.value)).String()
		if actual != c.expected {
			t.Errorf("String(%v): expected %s, actual %s", c.value, c.expected, actual)
		}
	}
}

func TestStringValue(t *testing.T) {
	cases := []struct {
		value    *string
		expected string
	}{
		{
			nil,
			"string",
		},
		{
			String("string"),
			"string",
		},
	}

	for _, c := range cases {
		actual := reflect.TypeOf(StringValue(c.value)).String()
		if actual != c.expected {
			t.Errorf("StringValue(%v): expected %s, actual %s", c.value, c.expected, actual)
		}
	}
}

func TestInt(t *testing.T) {
	cases := []struct {
		value    int
		expected string
	}{
		{
			123,
			"*int",
		},
	}

	for _, c := range cases {
		actual := reflect.TypeOf(Int(c.value)).String()
		if actual != c.expected {
			t.Errorf("Int(%v): expected %s, actual %s", c.value, c.expected, actual)
		}
	}
}

func TestIntValue(t *testing.T) {
	cases := []struct {
		value    *int
		expected string
	}{
		{
			nil,
			"int",
		},
		{
			Int(123),
			"int",
		},
	}

	for _, c := range cases {
		actual := reflect.TypeOf(IntValue(c.value)).String()
		if actual != c.expected {
			t.Errorf("IntValue(%v): expected %s, actual %s", c.value, c.expected, actual)
		}
	}
}
