package bencode

import (
	"errors"
	"strconv"
	"strings"
)

// 5:hello -> "hello"
func DecodeString(bencodedString string) (string, int, error) {
	if len(bencodedString) == 0 || bencodedString[0] < '0' || bencodedString[0] > '9' {
		return "", 0, errors.New("not a string")
	}

	colonIndex := strings.IndexByte(bencodedString, ':')

	stringLength, err := strconv.Atoi(bencodedString[:colonIndex])
	if err != nil {
		return "", 0, err
	}

	start := colonIndex + 1
	end := start + stringLength
	if end > len(bencodedString) {
		return "", 0, errors.New("string too short for declared length")
	}

	return bencodedString[start:end], end, nil
}

// i23e -> 23
func DecodeInteger(bencodedString string) (int, int, error) {
	if string(bencodedString[0]) != "i" {
		return 0, 0, errors.New("not an integer")
	}

	endIndex := strings.IndexByte(bencodedString, 'e')

	value := bencodedString[1:endIndex]
	res, err := strconv.Atoi(value)
	if err != nil {
		return 0, 0, err
	}

	totalConsumed := len(value) + 2
	return res, totalConsumed, nil
}

// l4:spam4:eggse -> []interface{}{"spam", "eggs"}
func DecodeList(bencodedString string) ([]interface{}, int, error) {
	if bencodedString[0] != 'l' {
		return nil, 0, errors.New("not a list")
	}
	var res []interface{}
	pos := 1

	for pos < len(bencodedString) && bencodedString[pos] != 'e' {
		for pos < len(bencodedString) {
			if bencodedString[pos] == 'e' {
				break
			}
			if pos >= len(bencodedString) {
				return nil, 0, errors.New("unexpected end of input")
			}
		}
		currentItem := bencodedString[pos]

		if bencodedString[pos] >= '0' && bencodedString[pos] <= '9' {
			value, consumed, err := DecodeString(bencodedString[pos:])
			if err != nil {
				return nil, 0, err
			}
			res = append(res, value)
			pos += consumed
		} else if currentItem == 'i' {
			value, consumed, err := DecodeInteger(bencodedString[pos:])
			if err != nil {
				return nil, 0, err
			}
			res = append(res, value)
			pos += consumed
		} else if currentItem == 'l' {
			value, consumed, err := DecodeList(bencodedString[pos:])
			if err != nil {
				return nil, 0, err
			}
			res = append(res, value)
			pos += consumed
		} else if currentItem == 'd' {
			// TODO DICTS
			return nil, 0, errors.New("dictionary decoding not yet implemented")
		} else {
			return nil, 0, errors.New("invalid bencode format")
		}

	}

	if pos >= len(bencodedString) || bencodedString[pos] != 'e' {
		return nil, 0, errors.New("no list end terminator")
	}
	return res, pos + 1, nil
}
