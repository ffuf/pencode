package pencode

import "encoding/json"

type JSONDecoder struct{}

type JSONInput struct {
	Input string `json:"input"`
}

func (u JSONDecoder) Encode(input []byte) ([]byte, error) {
	inputJson := `{"input":"` + string(input) + `"}`

	var out JSONInput
	if err := json.Unmarshal([]byte(inputJson), &out); err != nil {
		return []byte{}, err
	}

	return []byte(out.Input), nil
}

func (u JSONDecoder) HelpText() string {
	return "JSON decode"
}
