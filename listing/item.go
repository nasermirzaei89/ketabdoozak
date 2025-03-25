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
	Type  ItemContactInfoType `json:"type"`
	Value string              `json:"value"`
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
	ID            string            `json:"id"`
	Title         string            `json:"title"`
	OwnerID       string            `json:"ownerId"`
	OwnerName     string            `json:"ownerName"`
	LocationID    string            `json:"locationId"`
	LocationTitle string            `json:"locationTitle"`
	Types         []ItemType        `json:"types"`
	ContactInfo   []ItemContactInfo `json:"contactInfo"`
	Description   template.HTML     `json:"description"   swaggertype:"string"`
	Status        ItemStatus        `json:"status"`
	Lent          bool              `json:"lent"`
	ThumbnailURL  string            `json:"thumbnailUrl"`
	CreatedAt     time.Time         `json:"createdAt"`
	UpdatedAt     time.Time         `json:"updatedAt"`
	PublishedAt   *time.Time        `json:"publishedAt"`
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
