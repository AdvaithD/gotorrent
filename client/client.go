package client

import (
	"bytes"
	"fmt"
	"net"
	"time"
	"github.com/veggiedefender/torrent-client/bitfield"
	"github.com/veggiedefender/torrent-client/peers"
	"github.com/veggiedefender/torrent-client/message"
	"github.com/veggiedefender/torrent-client/handshake"
)

// tcp connection with a pier
type Client struct {
	Conn net.conn
	Choked bool
	Bitfield bitfield.Bitfield
	peer peers.Peer
	infoHash [20]byte
	peerId [20]byte
}

func completeHandshake(conn net.Conn, infohash, peerID [20]byte) (*handshake.Handshake, error) {
	conn.setDeadline(time.Now().Add(3 * time.Second))
	defer conn.SetDeadline(time.Time{}) // disable deadline
	
	req := handshake.New(infohash, peerID)
	_, err := conn.Write(req.Serealize())
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(res.InfoHash[:], infohash[:]) {
		return nil, fmt.Errorf("Expected infohash %x got %x", res.InfoHash, infohash)
	}
	return res,nil
}