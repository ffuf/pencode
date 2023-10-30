package pencode

import "html"

type HTMLEscaper struct{}

func (u HTMLEscaper) Encode(input []byte) ([]byte, error) {
	out := html.EscapeString(string(input))

	return []byte(out), nil
}

func (u HTMLEscaper) HelpText() string {
	return "HTML escape"
}

func (u HTMLEscaper) Type() string {
	return "encoders"
}
