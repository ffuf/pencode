package pencode

import (
	"crypto/sha512"
	"fmt"
)

type SHA384Hasher struct{}

func (m SHA384Hasher) Encode(input []byte) ([]byte, error) {
	return []byte(fmt.Sprintf("%x", sha512.Sum384(input))), nil
}

func (m SHA384Hasher) HelpText() string {
	return "SHA384 checksum"
}

func (m SHA384Hasher) Type() string {
	return "hashes"
}
