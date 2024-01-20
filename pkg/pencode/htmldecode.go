package pencode

import "html"

type HTMLUnescaper struct{}

func (u HTMLUnescaper) Encode(input []byte) ([]byte, error) {
	out := html.UnescapeString(string(input))

	return []byte(out), nil
}

func (u HTMLUnescaper) HelpText() string {
	return "HTML unescape"
}

func (u HTMLUnescaper) Type() string {
	return "decoders"
}
