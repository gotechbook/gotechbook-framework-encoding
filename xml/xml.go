package xml

import (
	"encoding/xml"
	encoding "github.com/gotechbook/gotechbook-framework-encoding"
)

const Name = "xml"

type codec struct {
}

func (e codec) Marshal(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

func (e codec) Unmarshal(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}

func (e codec) Name() string {
	return Name
}

func init() {
	encoding.RegisterEncoding(codec{})
}
