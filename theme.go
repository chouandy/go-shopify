package goshopify

import (
	"fmt"
	"time"
)

const themesBasePath = "admin/themes"

// ThemeAPI is an interface for interfacing with the themes endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/theme
type ThemeAPI interface {
	List(interface{}) ([]Theme, error)
}

// ThemeAPIOp handles communication with the theme related methods of
// the Shopify API.
type ThemeAPIOp struct {
	client *Client
}

// ThemeListOptions options for theme list
type ThemeListOptions struct {
	ListOptions
	Role string `url:"role,omitempty"`
}

// Theme represents a Shopify theme
type Theme struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Previewable  bool       `json:"previewable"`
	Processing   bool       `json:"processing"`
	Role         string     `json:"role"`
	ThemeStoreID int        `json:"theme_store_id"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

// ThemesResource is the result from the themes.json endpoint
type ThemesResource struct {
	Themes []Theme `json:"themes"`
}

// List all themes
func (s *ThemeAPIOp) List(options interface{}) ([]Theme, error) {
	path := fmt.Sprintf("%s.json", themesBasePath)
	resource := new(ThemesResource)
	err := s.client.Get(path, resource, options)
	return resource.Themes, err
}
