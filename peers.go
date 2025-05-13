package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jackpal/bencode-go"
)

type Response struct {
	Interval int    `bencode:"interval"`
	Peers    string `bencode:"peers"`
}

func DiscoverPeers(torrent TorrentFile) (string, error) {
	params := url.Values{}
	params.Add("info_hash", string(torrent.InfoHash))
	params.Add("peer_id", "radclientwritteningo")
	params.Add("port", "6881")
	params.Add("uploaded", "0")
	params.Add("downloaded", "0")
	params.Add("left", strconv.Itoa(torrent.Length))
	params.Add("compact", "1")

	fullURL := fmt.Sprintf("%s?%s", torrent.Announce, params.Encode())

	resp, err := http.Get(fullURL)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var res Response
	r := bytes.NewReader(body)
	err = bencode.Unmarshal(r, &res)
	if err != nil {
		return "", err
	}

	peers := ParsePeers(res.Peers)

	return peers, err
}

// get bytes from peers string
// loop over in chunks of 6
// first 4 = ip
// last 2 = port
func ParsePeers(peers string) string {
	b := []byte(peers)
	if len(b)%6 != 0 {
		return ""
	}

	numberOfChunks := len(b) / 6
	var result string

	for i := 0; i < numberOfChunks; i++ {
		offset := i * 6
		chunk := b[offset : offset+6]

		ip := net.IP(chunk[0:4])
		port := binary.BigEndian.Uint16(chunk[4:6])

		result += fmt.Sprintf("%s:%d\n", ip.String(), port)
	}

	return result
}
