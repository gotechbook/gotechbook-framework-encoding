package proto

import (
	encoding "github.com/gotechbook/gotechbook-framework-encoding"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	c := new(codec)
	if !reflect.DeepEqual(c.Name(), "proto") {
		t.Errorf("no expect float_key value: %v, but got: %v", c.Name(), "proto")
	}
}

func TestCodec(t *testing.T) {
	c := new(codec)
	model := encoding.TestModel{
		Id:    1,
		Name:  "go-tech-book-framework-encoding",
		Hobby: []string{"study", "eat", "play"},
	}

	m, err := c.Marshal(&model)
	if err != nil {
		t.Errorf("Marshal() should be nil, but got %s", err)
	}

	var res encoding.TestModel

	err = c.Unmarshal(m, &res)
	if err != nil {
		t.Errorf("Unmarshal() should be nil, but got %s", err)
	}
	if !reflect.DeepEqual(res.Id, model.Id) {
		t.Errorf("ID should be %v, but got %v", res.Id, model.Id)
	}
	if !reflect.DeepEqual(res.Name, model.Name) {
		t.Errorf("Name should be %s, but got %s", res.Name, model.Name)
	}
	if !reflect.DeepEqual(res.Hobby, model.Hobby) {
		t.Errorf("Hobby should be %s, but got %s", res.Hobby, model.Hobby)
	}
}
