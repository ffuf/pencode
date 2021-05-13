package pencode

import (
	"strings"
)

type StrToUpper struct{}
type StrToLower struct{}

func (s StrToUpper) Encode(input []byte) ([]byte, error) {
	return []byte(strings.ToUpper(string(input))), nil
}

func (s StrToUpper) HelpText() string {
	return "Convert string to uppercase"
}

func (s StrToUpper) Type() string {
	return "other"
}

func (s StrToLower) Encode(input []byte) ([]byte, error) {
	return []byte(strings.ToLower(string(input))), nil
}

func (s StrToLower) HelpText() string {
	return "Convert string to lowercase"
}

func (s StrToLower) Type() string {
	return "other"
}
