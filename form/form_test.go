package form

import (
	encoding "github.com/gotechbook/gotechbook-framework-encoding"
	"reflect"
	"testing"
)

type TestRequest struct {
	Param1 string `json:"param1,omitempty"`
	Param2 string `json:"param2,omitempty"`
}

type TestModel struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func TestFormCodecMarshal(t *testing.T) {
	req := &TestRequest{
		Param1: "param1",
		Param2: "param2",
	}
	data, err := encoding.GetEncoding(Name).Marshal(req)
	if err != nil {
		t.Errorf("marshal error: %v", err)
	}
	if !reflect.DeepEqual([]byte("param1=param1&param2=param2"), data) {
		t.Errorf("expect %v, got %v", []byte("param1=param1&param2=param2"), data)
	}
	m := &TestModel{
		ID:   1,
		Name: "go-tech-book-framework-encoding",
	}
	data, err = encoding.GetEncoding(Name).Marshal(m)
	t.Log(string(data))
	if err != nil {
		t.Errorf("expect %v,got %v", nil, err)
	}
	if !reflect.DeepEqual([]byte("id=1&name=go-tech-book-framework-encoding"), data) {
		t.Errorf("expect %v,got %v", []byte("id=1&name=go-tech-book-framework-encoding"), data)
	}
}
