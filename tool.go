package main

import (
	"errors"
	"fmt"
)

func HexToBytes(str string) ([]byte, error) {
	bytes := make([]byte, len(str)/2)
	var highBits byte
	var lowBits byte
	for i := 0; i < len(str); i += 2 {
		highBits = 0x00
		lowBits = 0x00
		switch {
		case str[i] >= '0' && str[i] <= '9':
			highBits = str[i] - '0'
		case str[i] >= 'a' && str[i] <= 'z':
			highBits = str[i] - 'a' + 10
		case str[i] >= 'A' && str[i] <= 'Z':
			highBits = str[i] - 'A' + 10
		default:
			return nil, errors.New(fmt.Sprintf("invalid hex character: %c", str[i]))
		}
		switch {
		case str[i+1] >= '0' && str[i] <= '9':
			lowBits = str[i+1] - '0'
		case str[i+1] >= 'a' && str[i] <= 'z':
			lowBits = str[i+1] - 'a' + 10
		case str[i+1] >= 'A' && str[i] <= 'Z':
			lowBits = str[i+1] - 'A' + 10
		default:
			return nil, errors.New(fmt.Sprintf("invalid hex character: %c", str[i]))
		}
		bytes[i/2] = highBits<<4 | lowBits

	}
	return bytes, nil
}

func main() {
	strA := fmt.Sprintf("%X", "\x12\x34\x56\x78\x90\xAB\xCD\xEF")
	fmt.Println(strA)
	strH, _ := HexToBytes(strA)
	fmt.Println(fmt.Sprintf("%X", strH))
	fmt.Printf("%T %V", strH, strH)
}
