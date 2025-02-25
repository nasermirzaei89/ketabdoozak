package www

import (
	"github.com/a-h/templ"
	"github.com/nasermirzaei89/ketabdoozak/listing"
	"github.com/nasermirzaei89/ketabdoozak/www/templates"
	"net/http"
	"time"
)

func (h *Handler) editItemPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		timeNow := time.Now()

		item := &listing.Item{
			ID:            "item-1",
			Title:         "آموزش برنامه نویسی C++ به زبان ساده",
			OwnerID:       "user1",
			OwnerName:     "ناصر میرزائی",
			LocationID:    "mashhad",
			LocationTitle: "مشهد",
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
<b>کتاب</b> آموزش برنامه نویسی سی پلاس، تقریبا نو. رایگان برای هرکسی که علاقه به خوندنش داره.
<br>
لطفا برای دریافت آن تماس بگیرید.
</p>
`,
			Status:       listing.ItemStatusPublished,
			Lent:         true,
			ThumbnailURL: "https://placehold.co/300x300",
			CreatedAt:    timeNow,
			UpdatedAt:    timeNow,
			PublishedAt:  &timeNow,
		}

		templ.Handler(templates.HTML(templates.EditItemPage(item))).ServeHTTP(w, r)
	}
}
