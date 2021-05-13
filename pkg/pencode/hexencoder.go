package pencode

import (
	"encoding/hex"
)

type HexEncoder struct{}

func (h HexEncoder) Encode(input []byte) ([]byte, error) {
	return []byte(hex.EncodeToString(input)), nil
}

func (h HexEncoder) HelpText() string {
	return "Hex string encoder"
}

func (h HexEncoder) Type() string {
	return "encoders"
}
