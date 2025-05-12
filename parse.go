package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/jackpal/bencode-go"
)

type bencodeTorrent struct {
	Announce string             `bencode:"announce"`
	Info     bencodeTorrentInfo `bencode:"info"`
}

type bencodeTorrentInfo struct {
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`
}

type TorrentFile struct {
	Announce    string
	InfoHash    []byte
	PieceHashes [][]byte
	PieceLength int
	Length      int
	Name        string
}

func ParseTorrentFile(file *os.File) (bencodeTorrent, error) {
	var res bencodeTorrent
	err := bencode.Unmarshal(file, &res)
	if err != nil {
		return bencodeTorrent{}, err
	}
	return res, nil
}

func toTorrentFile(torrent bencodeTorrent) (TorrentFile, error) {
	infoHash, err := getHash(torrent.Info)
	if err != nil {
		return TorrentFile{}, err
	}

	pieces := []byte(torrent.Info.Pieces)
	if len(pieces)%20 != 0 {
		return TorrentFile{}, fmt.Errorf("invalid pieces length")
	}

	var pieceHashes [][]byte
	for i := 0; i < len(pieces); i += 20 {
		pieceHashes = append(pieceHashes, pieces[i:i+20])
	}

	return TorrentFile{
		Announce:    torrent.Announce,
		InfoHash:    infoHash,
		PieceHashes: pieceHashes,
		PieceLength: torrent.Info.PieceLength,
		Length:      torrent.Info.Length,
		Name:        torrent.Info.Name,
	}, nil
}

func getHash(v any) ([]byte, error) {
	var buf bytes.Buffer
	err := bencode.Marshal(&buf, v)
	if err != nil {
		return nil, err
	}
	hash := sha1.Sum(buf.Bytes())
	return hash[:], nil
}
