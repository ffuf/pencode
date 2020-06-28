# pencode - complex payload encoder

Pencode is a tool that helps you to create payload encoding chains. It has been designed to be used in automation whereever
it is required to apply multiple encodings to a payload (and possibly inserting the payload to a template in between).

`pencoder` can be used as a standalone command line tool or as a library for other Go programs.


## Installation
```
go get -u github.com/ffuf/pencode/cmd/pencode
```

### Usage

```
pencode - complex payload encoder v0.1

Usage: pencode ENCODER1 ENCODER2 ENCODER3...

pencode reads input from stdin, which is typically piped from another process.

Available encoders:
  b64decode         - Base64 decoder
  b64encode         - Base64 encoder
  hexdecode         - Hex string decoder
  hexencode         - Hex string encoder
  unicodedecode     - Unicode escape string decode
  unicodeencodeall  - Unicode escape string encode (all characters)
  urldecode         - URL decode
  urlencode         - URL encode reserved characters
  urlencodeall      - URL encode all characters
  utf16             - UTF-16 encoder (Little Endian)
  utf16be           - UTF-16 encoder (Little Endian)
```

To urlencode, base64encode and hex encode a string:

```
$ echo 'what%ever'|pencode urlencode b64encode hexencode
64326868644355794e5756325a58493d
```

### Available encoders

- Base64 encoder
- Base64 decoder
- Hex string encoder
- Hex string decoder
- Unicode escape string encoder
- Unicode escape string decoder
- URL encoder (reserved characters)
- URL encoder (all characters)
- URL decoder
- UTF16 encoder (Little Endian)
- UTF16 encoder (Big Endian)

### Upcoming

- Templating

### Usage as a library

```go
package main

import (
    "fmt"
    
    "github.com/ffuf/pencode/pkg/pencode"
)

func main() {
    inputdata := "Whatever you wish to run through the chain"
    # A slice of encoders in the preferred encoding chain execution order
    encoders := []string{
        "utf16",
        "b64encode",
    }
    chain := pencode.NewChain()
    err := chain.Initialize(encoders)
    if err != nil {
        panic(err)
    }
    output, err := chain.Encode([]byte(inputdata))
    if err != nil {
        panic(err)
    }
    fmt.Print(string(output))
}
```

## License

`pencode` is released under MIT license. See [LICENSE](https://github.com/ffuf/pencoder/blob/master/LICENSE).