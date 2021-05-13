package pencode

import (
	"encoding/hex"
)

type HexDecoder struct{}

func (h HexDecoder) Encode(input []byte) ([]byte, error) {
	cleaned := removeAllWhitespace(string(input))
	return hex.DecodeString(cleaned)
}

func (h HexDecoder) HelpText() string {
	return "Hex string decoder"
}

func (h HexDecoder) Type() string {
	return "decoders"
}
