package main

import "testing"

func TestTextMessage_Make(t *testing.T) {
	m := TextMessage{"Hello {{.foo}} from {{.fizz}}"}
	res, _ := m.Make(&map[string]string {"foo": "bar", "fizz": "buzz"})
	if res != "Hello bar from buzz" {
		t.Error("Got", res)
	}
}
