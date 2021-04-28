package pencode

import (
	"fmt"
	"sort"
)

var availableEncoders = map[string]Encoder{
	"b64encode":        Base64Encoder{},
	"b64decode":        Base64Decoder{},
	"hexencode":        HexEncoder{},
	"hexdecode":        HexDecoder{},
	"jsonescape":       JSONEscaper{},
	"jsonunescape":     JSONUnescaper{},
	"unicodedecode":    UnicodeDecode{},
	"unicodeencodeall": UnicodeEncodeAll{},
	"urlencode":        URLEncoder{},
	"urldecode":        URLDecoder{},
	"urlencodeall":     URLEncoderAll{},
	"utf16":            UTF16LEEncode{},
	"utf16be":          UTF16BEEncode{},
	"xmlescape":        XMLEscaper{},
	"xmlunescape":      XMLUnescaper{},
}

type Chain struct {
	Encoders    []Encoder
	initialized bool
	actions     []string
}

func NewChain() *Chain {
	c := Chain{initialized: false}
	return &c
}

//Initialize loops through requested names for encoders and sets up the encoder chain. If an unknown encoder is
//requested, error will be returned.
func (c *Chain) Initialize(actions []string) error {
	c.actions = actions
	c.Encoders = make([]Encoder, 0)
	for _, a := range actions {
		if c.HasEncoder(a) {
			c.Encoders = append(c.Encoders, availableEncoders[a])
		} else {
			return fmt.Errorf("Encoder %s requested but not found.\n", a)
		}
	}
	c.initialized = true
	return nil
}

func (c *Chain) Encode(input []byte) ([]byte, error) {
	var err error
	if !c.initialized {
		return []byte{}, fmt.Errorf("Encoder chain not initialized.\n")
	}
	for _, e := range c.Encoders {
		input, err = e.Encode(input)
		if err != nil {
			return []byte{}, err
		}
	}
	return input, nil
}

//HasEncoder returns true if encoder with a specified name is configured
func (c *Chain) HasEncoder(name string) bool {
	if _, ok := availableEncoders[name]; ok {
		return true
	}
	return false
}

func (c *Chain) GetEncoders() []string {
	// Sort the encoders alphabetically
	names := make([]string, 0, len(availableEncoders))
	for e := range availableEncoders {
		names = append(names, e)
	}
	sort.Strings(names)
	return names
}

//Usage prints the help string for each  configured encoder
func (c *Chain) Usage() {
	// Calculate maximum keyword length for nice help formatting
	max_length := 0
	for k := range availableEncoders {
		if len(k) > max_length {
			max_length = len(k)
		}
	}
	format := fmt.Sprintf("  %%-%ds- %%s\n", max_length+2)
	names := c.GetEncoders()

	for _, n := range names {
		v := availableEncoders[n]
		fmt.Printf(format, n, v.HelpText())
	}
}
