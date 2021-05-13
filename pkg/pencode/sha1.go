package pencode

import (
	"crypto/sha1"
	"fmt"
)

type SHA1Hasher struct{}

func (m SHA1Hasher) Encode(input []byte) ([]byte, error) {
	return []byte(fmt.Sprintf("%x", sha1.Sum(input))), nil
}

func (m SHA1Hasher) HelpText() string {
	return "SHA1 checksum"
}

func (m SHA1Hasher) Type() string {
	return "hashes"
}
