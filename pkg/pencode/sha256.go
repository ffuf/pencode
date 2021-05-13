package pencode

import (
	"crypto/sha256"
	"fmt"
)

type SHA256Hasher struct{}

func (m SHA256Hasher) Encode(input []byte) ([]byte, error) {
	return []byte(fmt.Sprintf("%x", sha256.Sum256(input))), nil
}

func (m SHA256Hasher) HelpText() string {
	return "SHA256 checksum"
}

func (m SHA256Hasher) Type() string {
	return "hashes"
}
