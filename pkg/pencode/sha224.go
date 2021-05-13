package pencode

import (
	"crypto/sha256"
	"fmt"
)

type SHA224Hasher struct{}

func (m SHA224Hasher) Encode(input []byte) ([]byte, error) {
	return []byte(fmt.Sprintf("%x", sha256.Sum224(input))), nil
}

func (m SHA224Hasher) HelpText() string {
	return "SHA224 checksum"
}

func (m SHA224Hasher) Type() string {
	return "hashes"
}
