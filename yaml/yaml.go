package yaml

import (
	encoding "github.com/gotechbook/gotechbook-framework-encoding"
	"gopkg.in/yaml.v3"
)

const Name = "yaml"

type codec struct {
}

func (e codec) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (e codec) Unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}

func (e codec) Name() string {
	return Name
}

func init() {
	encoding.RegisterEncoding(codec{})
}
