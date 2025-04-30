package bencode

import (
	"strconv"
	"strings"
)

func DecodeString(bencodedString string) (string, error) {
	if bencodedString == "0:" {
		return "", nil
	}

	stringValue := strings.Split(bencodedString, ":")
	stringLength, err := strconv.Atoi(string(stringValue[0]))
	if err != nil {
		return "", err
	}

	var result strings.Builder

	for i := 0; i < stringLength; i++ {
		result.WriteByte(stringValue[1][:stringLength][i])
	}
	return result.String(), nil
}
