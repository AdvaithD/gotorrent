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

func recvBitfield(conn net.Conn) (bitfield.Bitfield, error) {
	conn.SetDeadline(time.Now().Add(5 * time.Second))
	defer conn.SetDeadline(time.Time{})

	msg, err := message.Read(conn)
	if err != nil {
		return nil, err
	}
	if msg == nil {
		err := fmt.Errorf("Expected bitfield but got %s", msg)
		return nil,err
	}
	if msg.ID != message.MsgBitfield {
		err := fmt.Errorf("Expected bitfield but got IF %d", msg.ID)
		return nil, err
	}
	return msg.Payload, nil
}

func New(peer peers.Peer, peerID, infoHash [20]byte) (*Client, error) {
	conn, err := net.DialTimeout("tcp", peer.String(), 3 * time.Second)
	if err != nil {
		return nil, err
	}
	_, err = completeHandshake(conn, infohash, peerID)
	if err != nil {
		conn.Close()
		return nil, err
	}
	
	bf, err := recvBitfield(conn)
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &Client{
		Conn: conn,
		Choked: true,
		Birfield: bf,
		peer: peer,
		infoHash: infoHash,
		peerID: peerID,
	}, nil
}