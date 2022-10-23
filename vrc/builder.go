package vrc

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/m-oons/headpat-osc/database"
)

const maxMessageLength = 144

var args = map[string]func() string{
	"time12h": getTime12h,
	"time24h": getTime24h,
	"count":   getHeadpatCount,
	"plural":  getPlural,
}

func buildMessage(text string) string {
	for arg, fn := range args {
		old := fmt.Sprintf("{{%s}}", arg)
		if strings.Contains(text, old) {
			new := fn()
			text = strings.ReplaceAll(text, old, new)
		}
	}
	return ensureMessageLength(text)
}

func ensureMessageLength(text string) string {
	if len(text) <= maxMessageLength {
		return text
	}
	return text[:maxMessageLength]
}

func getTime12h() string {
	return time.Now().Format("3:04PM")
}

func getTime24h() string {
	return time.Now().Format("15:04")
}

func getHeadpatCount() string {
	count := database.GetHeadpatCount()
	return strconv.FormatInt(count, 10)
}

func getPlural() string {
	count := database.GetHeadpatCount()
	if count == 1 {
		return ""
	}
	return "s"
}
