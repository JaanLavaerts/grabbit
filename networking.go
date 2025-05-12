package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jackpal/bencode-go"
)

type Response struct {
	Interval int    `bencode:"interval"`
	Peers    string `bencode:"peers"`
}

func DiscoverPeers(torrent TorrentFile) (Response, error) {
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
		return Response{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	var res Response
	r := bytes.NewReader(body)
	err = bencode.Unmarshal(r, &res)
	if err != nil {
		return Response{}, err
	}
	return Response{Interval: res.Interval, Peers: res.Peers}, nil
}

//TODO PARSE PEERS FUNCTION
