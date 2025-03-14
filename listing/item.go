package listing

import (
	"context"
	"html/template"
	"time"
)

type ItemType string

const (
	ItemTypeDonate   ItemType = "donate"
	ItemTypeExchange ItemType = "exchange"
	ItemTypeLend     ItemType = "lend"
	ItemTypeSell     ItemType = "sell"
)

type ItemContactInfoType string

const (
	ItemContactInfoTypePhoneNumber ItemContactInfoType = "phoneNumber"
	ItemContactInfoTypeSMS         ItemContactInfoType = "sms"
	ItemContactInfoTypeTelegram    ItemContactInfoType = "telegram"
	ItemContactInfoTypeWhatsapp    ItemContactInfoType = "whatsapp"
)

type ItemContactInfo struct {
	Type  ItemContactInfoType
	Value string
}

type ItemStatus string

const (
	ItemStatusDraft         ItemStatus = "draft"
	ItemStatusPendingReview ItemStatus = "pendingReview"
	ItemStatusPublished     ItemStatus = "published"
	ItemStatusRejected      ItemStatus = "rejected"
	ItemStatusExpired       ItemStatus = "expired"
	ItemStatusArchived      ItemStatus = "archived"
	ItemStatusDeleted       ItemStatus = "deleted"
)

type Item struct {
	ID            string
	Title         string
	OwnerID       string
	OwnerName     string
	LocationID    string
	LocationTitle string
	Types         []ItemType
	ContactInfo   []ItemContactInfo
	Description   template.HTML
	Status        ItemStatus
	Lent          bool
	ThumbnailURL  string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	PublishedAt   *time.Time
}

type ListItemsRequest struct {
	Query   string
	OwnerID string
	Status  ItemStatus
}

type ItemRepository interface {
	List(ctx context.Context, req *ListItemsRequest) (items []*Item, err error)
	Get(ctx context.Context, itemID string) (item *Item, err error)
	Replace(ctx context.Context, id string, item *Item) (err error)
	Insert(ctx context.Context, item *Item) (err error)
}
