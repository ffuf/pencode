package pencode

import "encoding/xml"

type XMLUnescaper struct{}

type XMLInput struct {
	Input string `xml:"input"`
}

func (u XMLUnescaper) Encode(input []byte) ([]byte, error) {
	inputXML := `<xml><input>` + string(input) + `</input></xml>`

	var out XMLInput
	if err := xml.Unmarshal([]byte(inputXML), &out); err != nil {
		return []byte{}, err
	}

	return []byte(out.Input), nil
}

func (u XMLUnescaper) HelpText() string {
	return "XML unescape"
}

func (u XMLUnescaper) Type() string {
	return "decoders"
}
