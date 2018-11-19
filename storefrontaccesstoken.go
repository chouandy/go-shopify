package goshopify

import (
	"fmt"
	"time"
)

const storefrontAccessTokensBasePath = "admin/storefront_access_tokens"

// StorefrontAccessTokenAPI is an interface for interfacing with the storefront access
// token endpoints of the Shopify API.
// See: https://help.shopify.com/api/reference/access/storefrontaccesstoken
type StorefrontAccessTokenAPI interface {
	List(interface{}) ([]StorefrontAccessToken, error)
	Create(StorefrontAccessToken) (*StorefrontAccessToken, error)
	Delete(int) error
}

// StorefrontAccessTokenAPIOp handles communication with the storefront access token
// related methods of the Shopify API.
type StorefrontAccessTokenAPIOp struct {
	client *Client
}

// StorefrontAccessToken represents a Shopify storefront access token
type StorefrontAccessToken struct {
	ID                int        `json:"id,omitempty"`
	Title             string     `json:"title,omitempty"`
	AccessToken       string     `json:"access_token,omitempty"`
	AccessScope       string     `json:"access_scope,omitempty"`
	AdminGraphqlAPIID string     `json:"admin_graphql_api_id,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
}

// StorefrontAccessTokenResource represents the result from the admin/storefront_access_tokens.json endpoint
type StorefrontAccessTokenResource struct {
	StorefrontAccessToken *StorefrontAccessToken `json:"storefront_access_token"`
}

// StorefrontAccessTokensResource is the root object for a storefront access tokens get request.
type StorefrontAccessTokensResource struct {
	StorefrontAccessTokens []StorefrontAccessToken `json:"storefront_access_tokens"`
}

// List storefront access tokens
func (s *StorefrontAccessTokenAPIOp) List(options interface{}) ([]StorefrontAccessToken, error) {
	path := fmt.Sprintf("%s.json", storefrontAccessTokensBasePath)
	resource := new(StorefrontAccessTokensResource)
	err := s.client.Get(path, resource, options)
	return resource.StorefrontAccessTokens, err
}

// Create a new storefront access token
func (s *StorefrontAccessTokenAPIOp) Create(storefrontAccessToken StorefrontAccessToken) (*StorefrontAccessToken, error) {
	path := fmt.Sprintf("%s.json", storefrontAccessTokensBasePath)
	wrappedData := StorefrontAccessTokenResource{StorefrontAccessToken: &storefrontAccessToken}
	resource := new(StorefrontAccessTokenResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.StorefrontAccessToken, err
}

// Delete an existing storefront access token
func (s *StorefrontAccessTokenAPIOp) Delete(ID int) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", storefrontAccessTokensBasePath, ID))
}
