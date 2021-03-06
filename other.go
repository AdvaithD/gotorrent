import (
"github.com/jackpal/bencode-go"
)

type bencodeInfo struct {
  Pieces      string `bencode:"pieces"`
  PieceLength int `bencode:"piece length"`
  Length      int `bencode:"length"`
  Name        string `bencode:"name"`
}

type bencodeTorrent struct {
  Announce string      `bencode:"announce"`
  Info     bencodeInfo `bencode:"info"`
}


func Open(r io.Reader) (*bencodeTorrent, error) {
  bto := bencodeTorrent{}
  err := bencode.Unmarshal(r, &bto)
  if err != nul {
    return nil, err 
  }
  return &bto, nil
}
