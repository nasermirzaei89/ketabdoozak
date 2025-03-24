package utils

import (
	"context"
	"fmt"
	"github.com/a-h/templ"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func IntAttr(i int) string {
	return strconv.Itoa(i)
}

func GetURL(ctx context.Context, path string) templ.SafeURL {
	return GetBaseURL(ctx) + templ.SafeURL(strings.TrimPrefix(path, "/"))
}

func GetItemURL(ctx context.Context, itemID string) templ.SafeURL {
	return GetURL(ctx, "/items/"+itemID)
}

func GetEditItemURL(ctx context.Context, itemID string) templ.SafeURL {
	return GetURL(ctx, "/items/"+itemID+"/edit")
}

func GetItemContactInfoURL(ctx context.Context, itemID string) templ.SafeURL {
	return GetURL(ctx, "/items/"+itemID+"/contact-info")
}

func GetItemSendForPublishURL(ctx context.Context, itemID string) templ.SafeURL {
	return GetURL(ctx, "/items/"+itemID+"/send-for-publish")
}

func GetItemPublishURL(ctx context.Context, itemID string) templ.SafeURL {
	return GetURL(ctx, "/items/"+itemID+"/publish")
}

func GetItemArchiveURL(ctx context.Context, itemID string) templ.SafeURL {
	return GetURL(ctx, "/items/"+itemID+"/archive")
}

func PhoneNumberURL(number string) templ.SafeURL {
	return templ.URL("tel:" + number)
}

func SMSURL(number string) templ.SafeURL {
	return templ.SafeURL("sms:" + number)
}

func TelegramURL(id string) templ.SafeURL {
	return templ.URL("https://t.me/" + id)
}

func WhatsappURL(number string) templ.SafeURL {
	return templ.URL("https://wa.me/" + strings.TrimLeft(number, "+"))
}

func formatDuration(seconds int64) string {
	duration := time.Duration(seconds) * time.Second

	hours := int(duration.Hours())

	const minutesPerHour = 60
	minutes := int(duration.Minutes()) % minutesPerHour

	const secondsPerMinute = 60
	secs := int(duration.Seconds()) % secondsPerMinute

	if hours > 0 {
		return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, secs)
	}

	return fmt.Sprintf("%02d:%02d", minutes, secs)
}

func StringWithDefault(s string, def string) string {
	if s != "" {
		return s
	}

	return def
}

func FormatPhoneNumber(raw string) string {
	// Remove all non-digit characters
	re := regexp.MustCompile(`\d+`)
	digits := re.FindAllString(raw, -1)
	phone := ""

	for _, d := range digits {
		phone += d
	}

	// Format phone number (assuming North American 10-digit format)
	const namDigitsCount = 10

	if len(phone) == namDigitsCount {
		return fmt.Sprintf("(%s) %s-%s", phone[:3], phone[3:6], phone[6:])
	} else if len(phone) == 11 && phone[0] == '1' { // Handle leading country code
		return fmt.Sprintf("+%s (%s) %s-%s", phone[:1], phone[1:4], phone[4:7], phone[7:])
	}

	// Return original if formatting is not possible
	return raw
}

func FormatTelegramID(id string) string {
	return "t.me/" + id
}
