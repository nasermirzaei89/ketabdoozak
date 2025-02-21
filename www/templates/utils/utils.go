package utils

import (
	"context"
	"fmt"
	"github.com/a-h/templ"
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

func formatDuration(seconds uint64) string {
	duration := time.Duration(seconds) * time.Second
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	secs := int(duration.Seconds()) % 60

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
