package pencode

import "encoding/json"

type JSONUnescaper struct{}

type JSONInput struct {
	Input string `json:"input"`
}

func (u JSONUnescaper) Encode(input []byte) ([]byte, error) {
	inputJson := `{"input":"` + string(input) + `"}`

	var out JSONInput
	if err := json.Unmarshal([]byte(inputJson), &out); err != nil {
		return []byte{}, err
	}

	return []byte(out.Input), nil
}

func (u JSONUnescaper) HelpText() string {
	return "JSON unescape"
}

func (u JSONUnescaper) Type() string {
	return "decoders"
}
