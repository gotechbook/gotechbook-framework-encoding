package encoding

import (
	"encoding/xml"
	"fmt"
	"runtime/debug"
	"testing"
)

type codec struct{}

func (c codec) Marshal(v interface{}) ([]byte, error) {
	panic("implement me")
}

func (c codec) Unmarshal(data []byte, v interface{}) error {
	panic("implement me")
}

func (c codec) Name() string {
	return ""
}

// codec2 is a Codec implementation with xml.
type codec2 struct{}

func (codec2) Marshal(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

func (codec2) Unmarshal(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}

func (codec2) Name() string {
	return "xml"
}

func TestRegisterCodec(t *testing.T) {
	f := func() { RegisterEncoding(nil) }
	funcDidPanic, panicValue, _ := didPanic(f)
	if !funcDidPanic {
		t.Fatalf(fmt.Sprintf("func should panic\n\tPanic value:\t%#v", panicValue))
	}
	if panicValue != "cannot register a nil encoding" {
		t.Fatalf("panic error got %s want cannot register a nil encoding", panicValue)
	}
	f = func() {
		RegisterEncoding(codec{})
	}
	funcDidPanic, panicValue, _ = didPanic(f)
	if !funcDidPanic {
		t.Fatalf(fmt.Sprintf("func should panic\n\tPanic value:\t%#v", panicValue))
	}
	if panicValue != "register name is empty" {
		t.Fatalf("panic error got %s want register name is empty", panicValue)
	}
	codec := codec2{}
	RegisterEncoding(codec)
	got := GetEncoding("xml")
	if got != codec {
		t.Fatalf("RegisterCodec(%v) want %v got %v", codec, codec, got)
	}
}

// PanicTestFunc defines a func that should be passed to assert.Panics and assert.NotPanics
// methods, and represents a simple func that takes no arguments, and returns nothing.
type PanicTestFunc func()

// didPanic returns true if the function passed to it panics. Otherwise, it returns false.
func didPanic(f PanicTestFunc) (bool, interface{}, string) {
	didPanic := false
	var message interface{}
	var stack string
	func() {
		defer func() {
			if message = recover(); message != nil {
				didPanic = true
				stack = string(debug.Stack())
			}
		}()
		f()
	}()
	return didPanic, message, stack
}