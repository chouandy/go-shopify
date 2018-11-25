package goshopify

import (
	"reflect"
	"testing"

	"gopkg.in/jarcoal/httpmock.v1"
)

func productListingTests(t *testing.T, productListing ProductListing) {
	// Check that ID is assigned to the returned product listing
	expectedInt := 921728736
	if productListing.ProductID != expectedInt {
		t.Errorf("ProductListing.ProductID returned %+v, expected %+v", productListing.ProductID, expectedInt)
	}
}

func TestProductListingList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/product_listings.json",
		httpmock.NewStringResponder(200, `{"product_listings": [{"product_id":1},{"product_id":2}]}`))

	productListings, err := client.ProductListing.List(nil)
	if err != nil {
		t.Errorf("ProductListing.List returned error: %v", err)
	}

	expected := []ProductListing{{ProductID: 1}, {ProductID: 2}}
	if !reflect.DeepEqual(productListings, expected) {
		t.Errorf("ProductListing.List returned %+v, expected %+v", productListings, expected)
	}
}

func TestProductListingListFilterByIDs(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/product_listings.json?product_ids=1,2,3",
		httpmock.NewStringResponder(200, `{"product_listings": [{"product_id":1},{"product_id":2},{"product_id":3}]}`))

	listOptions := ProductListingListOptions{ProductIDs: []int{1, 2, 3}}

	productListings, err := client.ProductListing.List(listOptions)
	if err != nil {
		t.Errorf("ProductListing.List returned error: %v", err)
	}

	expected := []ProductListing{{ProductID: 1}, {ProductID: 2}, {ProductID: 3}}
	if !reflect.DeepEqual(productListings, expected) {
		t.Errorf("ProductListing.List returned %+v, expected %+v", productListings, expected)
	}
}

func TestProductListingListProductIDs(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/product_listings/product_ids.json",
		httpmock.NewStringResponder(200, `{"product_ids": [1,2]}`))

	productIDs, err := client.ProductListing.ListProductIDs(nil)
	if err != nil {
		t.Errorf("ProductListing.ListProductIDs returned error: %v", err)
	}

	expected := []int{1, 2}
	if !reflect.DeepEqual(productIDs, expected) {
		t.Errorf("ProductListing.ListProductIDs returned %+v, expected %+v", productIDs, expected)
	}
}

func TestProductListingCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/product_listings/count.json",
		httpmock.NewStringResponder(200, `{"count": 3}`))

	cnt, err := client.ProductListing.Count(nil)
	if err != nil {
		t.Errorf("ProductListing.Count returned error: %v", err)
	}

	expected := 3
	if cnt != expected {
		t.Errorf("ProductListing.Count returned %d, expected %d", cnt, expected)
	}
}

func TestProductListingGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/product_listings/1.json",
		httpmock.NewStringResponder(200, `{"product_listing": {"product_id":1}}`))

	productListing, err := client.ProductListing.Get(1, nil)
	if err != nil {
		t.Errorf("ProductListing.Get returned error: %v", err)
	}

	expected := &ProductListing{ProductID: 1}
	if !reflect.DeepEqual(productListing, expected) {
		t.Errorf("ProductListing.Get returned %+v, expected %+v", productListing, expected)
	}
}

func TestProductListingPublish(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", "https://fooshop.myshopify.com/admin/product_listings/1.json",
		httpmock.NewBytesResponder(200, loadFixture("product_listing.json")))

	returnedProductListing, err := client.ProductListing.Publish(1)
	if err != nil {
		t.Errorf("ProductListing.Publish returned error: %v", err)
	}

	productListingTests(t, *returnedProductListing)
}

func TestProductListingUnpublish(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", "https://fooshop.myshopify.com/admin/product_listings/1.json",
		httpmock.NewStringResponder(200, "{}"))

	err := client.ProductListing.Unpublish(1)
	if err != nil {
		t.Errorf("ProductListing.Unpublish returned error: %v", err)
	}
}
