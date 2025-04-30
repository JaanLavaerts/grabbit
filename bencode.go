package main

import (
	"strconv"
	"strings"
)

// 5:hello -> "hello"
// 10:hello jaan -> "hello jaan"
// 10:hello:jaan -> "hello:jaan"
func decodeString(bencodedString string) (string, error) {
	if bencodedString == "0:" {
		return "", nil
	}

	length, err := strconv.Atoi(string(strings.Split(bencodedString, ":")[0]))
	if err != nil {
		return "", err
	}

	var result strings.Builder

	for i := 1; i < length+1; i++ {
		result.WriteByte(bencodedString[i+2])
	}
	return result.String(), nil
}
