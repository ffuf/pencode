package pencode

import (
	"bytes"
)

const upperhex = "0123456789ABCDEF"

type URLEncoderAll struct{}

func (u URLEncoderAll) Encode(input []byte) ([]byte, error) {
	var buf bytes.Buffer
	for _, c := range input {
		// Add "%"
		buf.WriteByte(37)
		buf.WriteByte(upperhex[c>>4])
		buf.WriteByte(upperhex[c&15])
	}
	return buf.Bytes(), nil
}

func (u URLEncoderAll) HelpText() string {
	return "URL encode all characters"
}

func (u URLEncoderAll) Type() string {
	return "encoders"
}
