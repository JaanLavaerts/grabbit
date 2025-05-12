package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	command := os.Args[1]
	fileName := os.Args[2]

	if command == "info" {
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
		fmt.Println(res.InfoHash)
	}

}
