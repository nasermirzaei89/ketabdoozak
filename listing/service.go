package listing

import (
	"context"
	"html/template"
)

const ServiceName = "github.com/nasermirzaei89/ketabdoozak/listing"

type Service interface {
	ListLocations(ctx context.Context) (rsp *ListLocationsResponse, err error)
	ListPublishedItems(ctx context.Context, q string) (rsp *ListItemsResponse, err error)
	ListMyItems(ctx context.Context, q string) (rsp *ListItemsResponse, err error)
	GetItem(ctx context.Context, itemID string) (item *Item, err error)
	GetPublishedItem(ctx context.Context, itemID string) (item *Item, err error)
	SendItemForPublish(ctx context.Context, itemID string) (err error)
	PublishItem(ctx context.Context, itemID string) (err error)
	ArchiveItem(ctx context.Context, itemID string) (err error)
	DeleteItem(ctx context.Context, itemID string) (err error)
	CreateItem(ctx context.Context, req *CreateItemRequest) (item *Item, err error)
	UpdateItem(ctx context.Context, itemID string, req *UpdateItemRequest) (err error)
}

type ListLocationsResponse struct {
	Items []*Location `json:"items"`
}

type ListItemsResponse struct {
	Items []*Item `json:"items"`
}

type CreateItemRequest struct {
	Title        string
	OwnerName    string
	LocationID   string
	Types        []ItemType
	ContactInfo  []ItemContactInfo
	Description  template.HTML
	ThumbnailURL string
	AsDraft      bool
}

type UpdateItemRequest struct {
	Title        string
	OwnerName    string
	LocationID   string
	Types        []ItemType
	ContactInfo  []ItemContactInfo
	Description  template.HTML
	Lent         bool
	ThumbnailURL string
	AsDraft      bool
}
