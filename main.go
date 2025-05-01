package main

import (
	"fmt"

	"github.com/JaanLavaerts/grabbit/bencode"
)

func main() {
	fmt.Println(bencode.DecodeList("ll4:nest4:nestl4:deep3:deeeeee"))
}
