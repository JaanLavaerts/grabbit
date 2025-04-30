package test

import (
	"testing"

	"github.com/JaanLavaerts/grabbit/bencode"
)

func TestStringDecode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"4:spam", "spam"},
		{"0:", ""},
		{"3:dog", "dog"},
		{"5:hello", "hello"},
		{"11:hello world", "hello world"},
		{"6:foobar", "foobar"},
		{"8:us:rname", "us:rname"},
	}

	for _, tt := range tests {
		result, err := bencode.DecodeString(tt.input)
		if err != nil {
			t.Errorf("decodeString(%q) returned error: %v", tt.input, err)
			continue
		}
		if result != tt.expected {
			t.Errorf("decodeString(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestIntegerDecode(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"i42e", 42},
		{"i0e", 0},
		{"i-1e", -1},
	}

	for _, tt := range tests {
		result, err := bencode.DecodeInteger(tt.input)
		if err != nil {
			t.Errorf("decodeInteger(%q) returned error: %v", tt.input, err)
			continue
		}
		if result != tt.expected {
			t.Errorf("decodeInteger(%q) = %d, want %d", tt.input, result, tt.expected)
		}
	}
}
