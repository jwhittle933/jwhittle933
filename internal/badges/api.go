package badges

import "fmt"

const (
	None = ""
	ColorPrimary = "011125"
	ColorSecondary = "172226"
	ColorAWS = "232F3E"
	ColorAzure = "0078D4"
	ColorHeroku = "430098"
)

func New(left, right, color, style, icon string) string {
	return fmt.Sprintf(
		"https://img.shields.io/badge/%s-%s-%s.svg?%slogo=%s",
		left,
		right,
		colorOrIgnore(color),
		styleOrNone(style),
		icon,
	)
}

func styleOrNone(s string) string {
	if s == "" {
		return s
	}

	return fmt.Sprintf("style=%s&", s)
}

func colorOrIgnore(s string) string {
	if s == "" {
		return "_"
	}

	return s
}
