package pencode

import (
	"bytes"
	"encoding/json"
)

type JSONEscaper struct{}

func (u JSONEscaper) Encode(input []byte) ([]byte, error) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false) // prevents encoding of < and > characters
	err := enc.Encode(string(input))
	if err != nil {
		return []byte{}, err
	}

	output := buf.Bytes()
	return output[1 : len(output)-2], nil
}

func (u JSONEscaper) HelpText() string {
	return "JSON escape"
}

func (u JSONEscaper) Type() string {
	return "encoders"
}
