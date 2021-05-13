package pencode

import (
	"encoding/base64"
)

type Base64Encoder struct{}

func (b Base64Encoder) Encode(input []byte) ([]byte, error) {
	output := base64.StdEncoding.EncodeToString(input)
	return []byte(output), nil
}

func (b Base64Encoder) HelpText() string {
	return "Base64 encoder"
}

func (b Base64Encoder) Type() string {
	return "encoders"
}
