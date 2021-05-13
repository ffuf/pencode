package pencode

type Encoder interface {
	Encode([]byte) ([]byte, error)
	HelpText() string
	Type() string
}
