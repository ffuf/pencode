package pencode

import (
	"bytes"
	"fmt"
)

type UnicodeEncodeAll struct{}

func (u UnicodeEncodeAll) Encode(input []byte) ([]byte, error) {
	var b bytes.Buffer
	runes := []rune(string(input))
	for _, r := range runes {
		b.WriteString("\\u")
		b.WriteString(fmt.Sprintf("%04x", int64(r)))
	}
	return b.Bytes(), nil
}

func (u UnicodeEncodeAll) HelpText() string {
	return "Unicode escape string encode (all characters)"
}

func (u UnicodeEncodeAll) Type() string {
	return "encoders"
}
