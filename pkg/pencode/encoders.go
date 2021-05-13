package pencode

import (
	"fmt"
	"sort"
	"strings"
)

var availableEncoders = map[string]Encoder{
	"b64encode":        Base64Encoder{},
	"b64decode":        Base64Decoder{},
	"filename.tmpl":    Template{},
	"hexencode":        HexEncoder{},
	"hexdecode":        HexDecoder{},
	"jsonescape":       JSONEscaper{},
	"jsonunescape":     JSONUnescaper{},
	"lower":            StrToLower{},
	"md5":              MD5Hasher{},
	"sha1":             SHA1Hasher{},
	"sha224":           SHA224Hasher{},
	"sha256":           SHA256Hasher{},
	"sha384":           SHA384Hasher{},
	"sha512":           SHA512Hasher{},
	"unicodedecode":    UnicodeDecode{},
	"unicodeencodeall": UnicodeEncodeAll{},
	"upper":            StrToUpper{},
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
		if ok, err := c.HasEncoder(a); ok {
			// Templates are a bit special
			if isTemplate(a) {
				tenc, err := NewTemplateEncoder(a)
				if err != nil {
					return err
				}
				c.Encoders = append(c.Encoders, tenc)
			} else {
				c.Encoders = append(c.Encoders, availableEncoders[a])
			}
		} else if err != nil {
			return err
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
func (c *Chain) HasEncoder(name string) (bool, error) {
	if _, ok := availableEncoders[name]; ok {
		return true, nil
	}
	// Check for template
	if isTemplate(name) {
		return hasTemplate(name)
	}
	return false, nil
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
	//names := c.GetEncoders()
	for _, t := range []string{"encoders", "decoders", "hashes", "other"} {
		fmt.Printf("%s\n", strings.ToUpper(t))
		list := []string{}
		for n, e := range availableEncoders {
			if e.Type() == t {
				list = append(list, fmt.Sprintf(format, n, e.HelpText()))
			}
		}
		sort.Strings(list)
		for _, i := range list {
			fmt.Print(i)
		}
		fmt.Println()
	}
}
