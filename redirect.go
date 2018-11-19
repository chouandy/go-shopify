package goshopify

import (
	"fmt"
)

const redirectsBasePath = "admin/redirects"

// RedirectAPI is an interface for interacting with the redirects
// endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/online_store/redirect
type RedirectAPI interface {
	List(interface{}) ([]Redirect, error)
	Count(interface{}) (int, error)
	Get(int, interface{}) (*Redirect, error)
	Create(Redirect) (*Redirect, error)
	Update(Redirect) (*Redirect, error)
	Delete(int) error
}

// RedirectAPIOp handles communication with the redirect related methods of the
// Shopify API.
type RedirectAPIOp struct {
	client *Client
}

// Redirect represents a Shopify redirect.
type Redirect struct {
	ID     int    `json:"id"`
	Path   string `json:"path"`
	Target string `json:"target"`
}

// RedirectResource represents the result from the redirects/X.json endpoint
type RedirectResource struct {
	Redirect *Redirect `json:"redirect"`
}

// RedirectsResource represents the result from the redirects.json endpoint
type RedirectsResource struct {
	Redirects []Redirect `json:"redirects"`
}

// List redirects
func (s *RedirectAPIOp) List(options interface{}) ([]Redirect, error) {
	path := fmt.Sprintf("%s.json", redirectsBasePath)
	resource := new(RedirectsResource)
	err := s.client.Get(path, resource, options)
	return resource.Redirects, err
}

// Count redirects
func (s *RedirectAPIOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", redirectsBasePath)
	return s.client.Count(path, options)
}

// Get individual redirect
func (s *RedirectAPIOp) Get(redirectID int, options interface{}) (*Redirect, error) {
	path := fmt.Sprintf("%s/%d.json", redirectsBasePath, redirectID)
	resource := new(RedirectResource)
	err := s.client.Get(path, resource, options)
	return resource.Redirect, err
}

// Create a new redirect
func (s *RedirectAPIOp) Create(redirect Redirect) (*Redirect, error) {
	path := fmt.Sprintf("%s.json", redirectsBasePath)
	wrappedData := RedirectResource{Redirect: &redirect}
	resource := new(RedirectResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Redirect, err
}

// Update an existing redirect
func (s *RedirectAPIOp) Update(redirect Redirect) (*Redirect, error) {
	path := fmt.Sprintf("%s/%d.json", redirectsBasePath, redirect.ID)
	wrappedData := RedirectResource{Redirect: &redirect}
	resource := new(RedirectResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Redirect, err
}

// Delete an existing redirect.
func (s *RedirectAPIOp) Delete(redirectID int) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", redirectsBasePath, redirectID))
}
