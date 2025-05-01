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
		result, _, err := bencode.DecodeString(tt.input)
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
		result, _, err := bencode.DecodeInteger(tt.input)
		if err != nil {
			t.Errorf("decodeInteger(%q) returned error: %v", tt.input, err)
			continue
		}
		if result != tt.expected {
			t.Errorf("decodeInteger(%q) = %d, want %d", tt.input, result, tt.expected)
		}
	}
}

func TestListDecode(t *testing.T) {
	tests := []struct {
		input    string
		expected []interface{}
	}{
		{"l4:spam4:eggse", []interface{}{"spam", "eggs"}},
		{"li1ei2ee", []interface{}{1, 2}},
		{"llee", []interface{}{[]interface{}{}}},
		{"ll4:foo4:baree", []interface{}{[]interface{}{"foo", "bar"}}},
		{"le", []interface{}{}},
	}

	for _, tt := range tests {
		got, _, err := bencode.DecodeList(tt.input)
		if err != nil {
			t.Errorf("DecodeList(%q) error: %v", tt.input, err)
			continue
		}
		if !equal(got, tt.expected) {
			t.Errorf("DecodeList(%q) = %v, want %v", tt.input, got, tt.expected)
		}
	}
}

func equal(a, b interface{}) bool {
	switch a := a.(type) {
	case []interface{}:
		bList, ok := b.([]interface{})
		if !ok || len(a) != len(bList) {
			return false
		}
		for i := range a {
			if !equal(a[i], bList[i]) {
				return false
			}
		}
		return true
	default:
		return a == b
	}
}
