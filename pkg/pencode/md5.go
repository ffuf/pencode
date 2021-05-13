package pencode

import (
	"crypto/md5"
	"fmt"
)

type MD5Hasher struct{}

func (m MD5Hasher) Encode(input []byte) ([]byte, error) {
	return []byte(fmt.Sprintf("%x", md5.Sum(input))), nil
}

func (m MD5Hasher) HelpText() string {
	return "MD5 sum"
}

func (m MD5Hasher) Type() string {
	return "hashes"
}
