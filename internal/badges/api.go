package badges

import "fmt"

const (
	None = ""
	ColorPrimary = "011125"
	ColorSecondary = "172226"
)

func New(left, right, color, icon string) string {
	return fmt.Sprintf(
		"https://img.shields.io/badge/%s%s%s.svg?logo=%s",
		valueOrDash(left),
		valueOrDash(right),
		valueOrDash(color),
		icon,
	)
}

func valueOrDash(s string) string {
	if s == "" {
		return "-"
	}

	return s
}

