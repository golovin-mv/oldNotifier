package oldNotifier

import (
	"testing"
)

func TestTextMessage_Make(t *testing.T) {
	m := NewTextMessage("Hello {{.foo}} from {{.fizz}}")
	m.AddParams(&map[string]string{"foo": "bar", "fizz": "buzz"})
	res, _ := m.Make()
	if res != "Hello bar from buzz" {
		t.Error("Got", res)
	}
}
