package pencode

import (
	"bytes"
	"encoding/binary"
	"unicode/utf16"
)

type UTF16LEEncode struct{}

func (u UTF16LEEncode) Encode(input []byte) ([]byte, error) {
	var b bytes.Buffer
	runes := []rune(string(input))
	utf16ints := utf16.Encode(runes)
	for _, r := range utf16ints {
		tmp := make([]byte, 2)
		binary.LittleEndian.PutUint16(tmp, r)
		b.Write(tmp)
	}
	return b.Bytes(), nil
}

func (u UTF16LEEncode) HelpText() string {
	return "UTF-16 encoder (Little Endian)"
}

func (u UTF16LEEncode) Type() string {
	return "encoders"
}
