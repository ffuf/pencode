package pencode

import (
	"bytes"
	"encoding/json"
)

type JSONEncoder struct{}

func (u JSONEncoder) Encode(input []byte) ([]byte, error) {
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

func (u JSONEncoder) HelpText() string {
	return "JSON encode reserved characters"
}
