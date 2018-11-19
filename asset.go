package goshopify

import (
	"fmt"
	"time"
)

const assetsBasePath = "admin/themes"

// AssetAPI is an interface for interfacing with the asset endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/asset
type AssetAPI interface {
	List(int, interface{}) ([]Asset, error)
	Get(int, string) (*Asset, error)
	Update(int, Asset) (*Asset, error)
	Delete(int, string) error
}

// AssetAPIOp handles communication with the asset related methods of
// the Shopify API.
type AssetAPIOp struct {
	client *Client
}

// Asset represents a Shopify asset
type Asset struct {
	Attachment  string     `json:"attachment"`
	ContentType string     `json:"content_type"`
	Key         string     `json:"key"`
	PublicURL   string     `json:"public_url"`
	Size        int        `json:"size"`
	SourceKey   string     `json:"source_key"`
	Src         string     `json:"src"`
	ThemeID     int        `json:"theme_id"`
	Value       string     `json:"value"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

// AssetResource is the result from the themes/x/assets.json?asset[key]= endpoint
type AssetResource struct {
	Asset *Asset `json:"asset"`
}

// AssetsResource is the result from the themes/x/assets.json endpoint
type AssetsResource struct {
	Assets []Asset `json:"assets"`
}

type assetGetOptions struct {
	Key     string `url:"asset[key]"`
	ThemeID int    `url:"theme_id"`
}

// List the metadata for all assets in the given theme
func (s *AssetAPIOp) List(themeID int, options interface{}) ([]Asset, error) {
	path := fmt.Sprintf("%s/%d/assets.json", assetsBasePath, themeID)
	resource := new(AssetsResource)
	err := s.client.Get(path, resource, options)
	return resource.Assets, err
}

// Get an asset by key from the given theme
func (s *AssetAPIOp) Get(themeID int, key string) (*Asset, error) {
	path := fmt.Sprintf("%s/%d/assets.json", assetsBasePath, themeID)
	options := assetGetOptions{
		Key:     key,
		ThemeID: themeID,
	}
	resource := new(AssetResource)
	err := s.client.Get(path, resource, options)
	return resource.Asset, err
}

// Update an asset
func (s *AssetAPIOp) Update(themeID int, asset Asset) (*Asset, error) {
	path := fmt.Sprintf("%s/%d/assets.json", assetsBasePath, themeID)
	wrappedData := AssetResource{Asset: &asset}
	resource := new(AssetResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Asset, err
}

// Delete an asset
func (s *AssetAPIOp) Delete(themeID int, key string) error {
	path := fmt.Sprintf("%s/%d/assets.json?asset[key]=%s", assetsBasePath, themeID, key)
	return s.client.Delete(path)
}
