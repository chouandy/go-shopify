package goshopify

import (
	"fmt"
	"time"
)

const pagesBasePath = "admin/pages"
const pagesResourceName = "pages"

// PageService is an interface for interacting with the pages
// endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/online_store/page
type PageService interface {
	List(interface{}) ([]Page, error)
	Count(interface{}) (int, error)
	Get(int, interface{}) (*Page, error)
	Create(Page) (*Page, error)
	Update(Page) (*Page, error)
	Delete(int) error

	// MetafieldsAPI used for Pages resource to communicate with Metafields
	// resource
	MetafieldsAPI
}

// PageServiceOp handles communication with the page related methods of the
// Shopify API.
type PageServiceOp struct {
	client *Client
}

// Page represents a Shopify page.
type Page struct {
	ID             int         `json:"id"`
	Author         string      `json:"author"`
	Handle         string      `json:"handle"`
	Title          string      `json:"title"`
	CreatedAt      *time.Time  `json:"created_at"`
	UpdatedAt      *time.Time  `json:"updated_at"`
	BodyHTML       string      `json:"body_html"`
	TemplateSuffix string      `json:"template_suffix"`
	PublishedAt    *time.Time  `json:"published_at"`
	ShopID         int         `json:"shop_id"`
	Metafields     []Metafield `json:"metafields"`
}

// PageResource represents the result from the pages/X.json endpoint
type PageResource struct {
	Page *Page `json:"page"`
}

// PagesResource represents the result from the pages.json endpoint
type PagesResource struct {
	Pages []Page `json:"pages"`
}

// List pages
func (s *PageServiceOp) List(options interface{}) ([]Page, error) {
	path := fmt.Sprintf("%s.json", pagesBasePath)
	resource := new(PagesResource)
	err := s.client.Get(path, resource, options)
	return resource.Pages, err
}

// Count pages
func (s *PageServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", pagesBasePath)
	return s.client.Count(path, options)
}

// Get individual page
func (s *PageServiceOp) Get(pageID int, options interface{}) (*Page, error) {
	path := fmt.Sprintf("%s/%d.json", pagesBasePath, pageID)
	resource := new(PageResource)
	err := s.client.Get(path, resource, options)
	return resource.Page, err
}

// Create a new page
func (s *PageServiceOp) Create(page Page) (*Page, error) {
	path := fmt.Sprintf("%s.json", pagesBasePath)
	wrappedData := PageResource{Page: &page}
	resource := new(PageResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Page, err
}

// Update an existing page
func (s *PageServiceOp) Update(page Page) (*Page, error) {
	path := fmt.Sprintf("%s/%d.json", pagesBasePath, page.ID)
	wrappedData := PageResource{Page: &page}
	resource := new(PageResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Page, err
}

// Delete an existing page.
func (s *PageServiceOp) Delete(pageID int) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", pagesBasePath, pageID))
}

// ListMetafields list metafields for a page
func (s *PageServiceOp) ListMetafields(pageID int, options interface{}) ([]Metafield, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: pagesResourceName, resourceID: pageID}
	return metafieldAPI.List(options)
}

// CountMetafields count metafields for a page
func (s *PageServiceOp) CountMetafields(pageID int, options interface{}) (int, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: pagesResourceName, resourceID: pageID}
	return metafieldAPI.Count(options)
}

// GetMetafield get individual metafield for a page
func (s *PageServiceOp) GetMetafield(pageID int, metafieldID int, options interface{}) (*Metafield, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: pagesResourceName, resourceID: pageID}
	return metafieldAPI.Get(metafieldID, options)
}

// CreateMetafield create a new metafield for a page
func (s *PageServiceOp) CreateMetafield(pageID int, metafield Metafield) (*Metafield, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: pagesResourceName, resourceID: pageID}
	return metafieldAPI.Create(metafield)
}

// UpdateMetafield update an existing metafield for a page
func (s *PageServiceOp) UpdateMetafield(pageID int, metafield Metafield) (*Metafield, error) {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: pagesResourceName, resourceID: pageID}
	return metafieldAPI.Update(metafield)
}

// DeleteMetafield delete an existing metafield for a page
func (s *PageServiceOp) DeleteMetafield(pageID int, metafieldID int) error {
	metafieldAPI := &MetafieldAPIOp{client: s.client, resource: pagesResourceName, resourceID: pageID}
	return metafieldAPI.Delete(metafieldID)
}
