package encoding

import (
	"strings"
)

var registeredEncodings = make(map[string]Encoding)

func RegisterEncoding(e Encoding) {
	if e == nil {
		panic("cannot register a nil encoding")
	}
	if e.Name() == "" {
		panic("register name is empty")
	}
	t := strings.ToLower(e.Name())
	registeredEncodings[t] = e
}

func GetEncoding(t string) Encoding {
	return registeredEncodings[t]
}
