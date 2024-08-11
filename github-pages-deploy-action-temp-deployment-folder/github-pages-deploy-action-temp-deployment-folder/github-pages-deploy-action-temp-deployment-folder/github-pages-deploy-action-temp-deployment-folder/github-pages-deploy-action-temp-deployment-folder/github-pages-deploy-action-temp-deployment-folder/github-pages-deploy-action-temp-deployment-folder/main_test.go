package main

import (
	"testing"
)

func TestSendMessage(t *testing.T) {
	expected := "success"
	result := SendMessage()

	if result != expected {
		t.Errorf("sendMessage() = %v; want %v", result, expected)
	}
}
