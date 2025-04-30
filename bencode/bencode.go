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

func DecodeInteger(bencodedString string) (int, error) {
	value := bencodedString[1 : len(bencodedString)-1]

	res, err := strconv.Atoi(value)
	if err != nil {
		return 0, nil
	}
	return res, nil
}
