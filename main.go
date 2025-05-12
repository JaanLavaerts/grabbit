package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	command := os.Args[1]
	fileName := os.Args[2]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	output, err := ParseTorrentFile(file)
	if err != nil {
		log.Fatal(err)
	}

	res, err := toTorrentFile(output)
	if err != nil {
		log.Fatal(err)
	}

	if command == "info" {
		fmt.Println(res)
	}

	if command == "peers" {
		output, _ := DiscoverPeers(res)
		fmt.Println(output)
	}

}
