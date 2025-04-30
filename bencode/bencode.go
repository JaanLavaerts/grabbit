package bencode

import (
	"strconv"
	"strings"
)

func DecodeString(bencodedString string) (string, error) {
	if bencodedString == "0:" {
		return "", nil
	}

	stringValue := strings.SplitN(bencodedString, ":", 2)
	stringLength, err := strconv.Atoi(stringValue[0])
	if err != nil {
		return "", err
	}

	return stringValue[1][:stringLength], nil
}

func DecodeInteger(bencodedString string) (int, error) {
	value := bencodedString[1 : len(bencodedString)-1]

	res, err := strconv.Atoi(value)
	if err != nil {
		return 0, nil
	}
	return res, nil
}
