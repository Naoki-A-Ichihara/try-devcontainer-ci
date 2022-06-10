package main

import "testing"

func TestHelloWorldStr(t *testing.T) {
	actual := HelloWorldStr()
	expected := "Hello, World!"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
