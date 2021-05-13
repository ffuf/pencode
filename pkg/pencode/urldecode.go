package pencode

import (
	"net/url"
)

type URLDecoder struct{}

func (u URLDecoder) Encode(input []byte) ([]byte, error) {
	output, err := url.QueryUnescape(string(input))
	if err != nil {
		return []byte{}, err
	}
	return []byte(output), err
}

func (u URLDecoder) HelpText() string {
	return "URL decode"
}

func (u URLDecoder) Type() string {
	return "decoders"
}
