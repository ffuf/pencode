package pencode

import (
	"strings"
	"unicode"
)

func removeAllWhitespace(input string) string {
	var b strings.Builder
	b.Grow(len(input))
	for _, ch := range input {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
