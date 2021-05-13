package pencode

import (
	"encoding/base64"
)

type Base64Decoder struct{}

func (b Base64Decoder) Encode(input []byte) ([]byte, error) {
	output, err := base64.StdEncoding.DecodeString(string(input))
	return output, err
}

func (b Base64Decoder) HelpText() string {
	return "Base64 decoder"
}

func (b Base64Decoder) Type() string {
	return "decoders"
}
