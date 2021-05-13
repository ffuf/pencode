package pencode

import (
	"io/ioutil"
	"os"
	"strings"
)

//isTemplate checks if the name ends with string ".tmpl"
func isTemplate(name string) bool {
	return strings.HasSuffix(name, ".tmpl")
}

//hasTemplate checks if the name points to an existing file
func hasTemplate(name string) (bool, error) {
	if _, err := os.Stat(name); err == nil {
		return true, nil
	} else {
		return false, err
	}
}

type Template struct {
	Data    []byte
	Keyword string
}

func NewTemplateEncoder(name string) (Encoder, error) {
	var ok bool
	var err error
	t := Template{}
	// Using static keyword for now
	t.Keyword = "#PAYLOAD#"

	if ok, err = hasTemplate(name); ok {
		t.Data, err = ioutil.ReadFile(name)
	}
	return t, err
}

func (t Template) Encode(input []byte) ([]byte, error) {
	return []byte(strings.ReplaceAll(string(t.Data), t.Keyword, string(input))), nil
}

func (t Template) HelpText() string {
	return "Replaces string #PAYLOAD# in content of a file that has .tmpl extension."
}

func (t Template) Type() string {
	return "other"
}
