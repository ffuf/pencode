package pencode

import (
	"bytes"
	"encoding/xml"
)

type XMLEscaper struct{}

func (u XMLEscaper) Encode(input []byte) ([]byte, error) {
	buf := &bytes.Buffer{}

	if err := xml.EscapeText(buf, input); err != nil {
		return []byte{}, err
	}

	output := buf.Bytes()
	return output, nil
}

func (u XMLEscaper) HelpText() string {
	return "XML escape"
}

func (u XMLEscaper) Type() string {
	return "encoders"
}
