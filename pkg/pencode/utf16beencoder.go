package pencode

import (
	"bytes"
	"encoding/binary"
	"unicode/utf16"
)

type UTF16BEEncode struct{}

func (u UTF16BEEncode) Encode(input []byte) ([]byte, error) {
	var b bytes.Buffer
	runes := []rune(string(input))
	utf16ints := utf16.Encode(runes)
	for _, r := range utf16ints {
		tmp := make([]byte, 2)
		binary.BigEndian.PutUint16(tmp, r)
		b.Write(tmp)
	}
	return b.Bytes(), nil
}

func (u UTF16BEEncode) HelpText() string {
	return "UTF-16 encoder (Big Endian)"
}

func (u UTF16BEEncode) Type() string {
	return "encoders"
}
