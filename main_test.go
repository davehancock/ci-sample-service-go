package main

import (
	"testing"
)

func TestCreateMessage(t *testing.T) {

	content := string(createMessage())

	if content !=  "Hello!" {
		t.Fail()
	}
}
