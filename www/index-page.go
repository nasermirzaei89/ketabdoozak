package www

import (
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/listing"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"net/http"
	"time"
)

func (h *Handler) indexPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			h.notFoundPageHandler()(w, r)

			return
		}

		timeNow := time.Now()

		items := []listing.Item{
			{
				ID:            "item-1",
				Title:         "آموزش برنامه نویسی C++ به زبان ساده",
				LocationID:    "tehran-saadatabad",
				LocationTitle: "تهران، سعادت آباد",
				Types: []listing.ItemType{
					listing.ItemTypeDonate,
					listing.ItemTypeExchange,
				},
				ContactInfo: []listing.ItemContactInfo{
					{
						Type:  listing.ItemContactInfoTypePhoneNumber,
						Value: "+1234567890",
					},
					{
						Type:  listing.ItemContactInfoTypeTelegram,
						Value: "t.me/user1",
					},
				},
				Description: `
<p>
کتاب آموزش برنامه نویسی سی پلاس، تقریبا نو. رایگان برای هرکسی که علاقه به خوندنش داره.
<br>
لطفا برای دریافت آن تماس بگیرید.
</p>
`,
				Status:       listing.ItemStatusPublished,
				Lent:         false,
				ThumbnailURL: "https://placehold.co/300x300",
				CreatedAt:    timeNow,
				UpdatedAt:    timeNow,
				PublishedAt:  &timeNow,
			},
			{
				ID:            "item-2",
				Title:         "آموزش گام به گام برنامه نویسی Go",
				LocationID:    "tehran-saadatabad",
				LocationTitle: "تهران، سعادت آباد",
				Types: []listing.ItemType{
					listing.ItemTypeDonate,
					listing.ItemTypeExchange,
				},
				ContactInfo: []listing.ItemContactInfo{
					{
						Type:  listing.ItemContactInfoTypePhoneNumber,
						Value: "+1234567890",
					},
					{
						Type:  listing.ItemContactInfoTypeTelegram,
						Value: "t.me/user1",
					},
				},
				Description: `
<p>
کتاب آموزش برنامه نویسی سی پلاس، تقریبا نو. رایگان برای هرکسی که علاقه به خوندنش داره.
<br>
لطفا برای دریافت آن تماس بگیرید.
</p>
`,
				Status:       listing.ItemStatusPublished,
				Lent:         false,
				ThumbnailURL: "https://placehold.co/300x300",
				CreatedAt:    timeNow,
				UpdatedAt:    timeNow,
				PublishedAt:  &timeNow,
			},
			{
				ID:            "item-3",
				Title:         "Dependency Injection in Go",
				LocationID:    "tehran-saadatabad",
				LocationTitle: "تهران، سعادت آباد",
				Types: []listing.ItemType{
					listing.ItemTypeDonate,
					listing.ItemTypeExchange,
				},
				ContactInfo: []listing.ItemContactInfo{
					{
						Type:  listing.ItemContactInfoTypePhoneNumber,
						Value: "+1234567890",
					},
					{
						Type:  listing.ItemContactInfoTypeTelegram,
						Value: "t.me/user1",
					},
				},
				Description: `
<p>
کتاب آموزش برنامه نویسی سی پلاس، تقریبا نو. رایگان برای هرکسی که علاقه به خوندنش داره.
<br>
لطفا برای دریافت آن تماس بگیرید.
</p>
`,
				Status:       listing.ItemStatusPublished,
				Lent:         false,
				ThumbnailURL: "https://placehold.co/300x300",
				CreatedAt:    timeNow,
				UpdatedAt:    timeNow,
				PublishedAt:  &timeNow,
			},
			{
				ID:            "item-4",
				Title:         "کتاب مدیریت پروژه",
				LocationID:    "tehran-saadatabad",
				LocationTitle: "تهران، سعادت آباد",
				Types: []listing.ItemType{
					listing.ItemTypeDonate,
					listing.ItemTypeExchange,
				},
				ContactInfo: []listing.ItemContactInfo{
					{
						Type:  listing.ItemContactInfoTypePhoneNumber,
						Value: "+1234567890",
					},
					{
						Type:  listing.ItemContactInfoTypeTelegram,
						Value: "t.me/user1",
					},
				},
				Description: `
<p>
کتاب آموزش برنامه نویسی سی پلاس، تقریبا نو. رایگان برای هرکسی که علاقه به خوندنش داره.
<br>
لطفا برای دریافت آن تماس بگیرید.
</p>
`,
				Status:       listing.ItemStatusPublished,
				Lent:         false,
				ThumbnailURL: "https://placehold.co/300x300",
				CreatedAt:    timeNow,
				UpdatedAt:    timeNow,
				PublishedAt:  &timeNow,
			},
			{
				ID:            "item-5",
				Title:         "Domain Driven Design by Eric Evans",
				LocationID:    "tehran-saadatabad",
				LocationTitle: "تهران، سعادت آباد",
				Types: []listing.ItemType{
					listing.ItemTypeDonate,
					listing.ItemTypeExchange,
				},
				ContactInfo: []listing.ItemContactInfo{
					{
						Type:  listing.ItemContactInfoTypePhoneNumber,
						Value: "+1234567890",
					},
					{
						Type:  listing.ItemContactInfoTypeTelegram,
						Value: "t.me/user1",
					},
				},
				Description: `
<p>
کتاب آموزش برنامه نویسی سی پلاس، تقریبا نو. رایگان برای هرکسی که علاقه به خوندنش داره.
<br>
لطفا برای دریافت آن تماس بگیرید.
</p>
`,
				Status:       listing.ItemStatusPublished,
				Lent:         false,
				ThumbnailURL: "https://placehold.co/300x300",
				CreatedAt:    timeNow,
				UpdatedAt:    timeNow,
				PublishedAt:  &timeNow,
			},
		}

		templ.Handler(templates.HTML(templates.IndexPage(items, r.URL.Query().Get("q")))).ServeHTTP(w, r)
	}
}
