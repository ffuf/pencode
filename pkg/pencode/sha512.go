package pencode

import (
	"crypto/sha512"
	"fmt"
)

type SHA512Hasher struct{}

func (m SHA512Hasher) Encode(input []byte) ([]byte, error) {
	return []byte(fmt.Sprintf("%x", sha512.Sum512(input))), nil
}

func (m SHA512Hasher) HelpText() string {
	return "SHA512 checksum"
}

func (m SHA512Hasher) Type() string {
	return "hashes"
}
