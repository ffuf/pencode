package pencode

import (
	"net/url"
)

type URLEncoder struct{}

func (u URLEncoder) Encode(input []byte) ([]byte, error) {
	return []byte(url.QueryEscape(string(input))), nil
}

func (u URLEncoder) HelpText() string {
	return "URL encode reserved characters"
}

func (u URLEncoder) Type() string {
	return "encoders"
}
