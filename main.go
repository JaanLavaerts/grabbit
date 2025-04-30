package main

import (
	"fmt"

	"github.com/JaanLavaerts/grabbit/bencode"
)

func main() {
	fmt.Println(bencode.DecodeString("5:hello"))
}
