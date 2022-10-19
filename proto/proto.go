package proto

import (
	encoding "github.com/gotechbook/gotechbook-framework-encoding"
	"google.golang.org/protobuf/proto"
)

const Name = "proto"

type codec struct {
}

func (e codec) Marshal(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (e codec) Unmarshal(data []byte, v interface{}) error {
	return proto.Unmarshal(data, v.(proto.Message))
}

func (e codec) Name() string {
	return Name
}

func init() {
	encoding.RegisterEncoding(codec{})
}
