package p2p

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"log"
	"runtime"
	"time"
	"github.com/AdvaithD/gotorrent/client"
	"github.com/AdvaithD/gotorrent/message"
	"github.com/AdvaithD/gotorrent/peers"
)

const MaxBlockSize = 16384
const MaxBacklog = 5

type Torrent struct {
	Peers []peers.Peer
	PeerID [20]byte
	InfoHash [20]byte
	PiecesHashes [][20]bytes
	PiecesLength int
	Length int
	Name String
}

type pieceWork struct {
	index int
	hash [20]byte
	length int
}

type pieceResult struct {
	index int
	buf []byte
}

type pieceProgress struct {
	index int
	client *client.Client
	bug []byte
	downloaded int
	requested int
	backlog int
}

func (state *pieceProgress) readMessage() {
	msg, err := state.client.Read()
	if err != nil {
		return err
	}
	switch msg.ID {
	case message.MsgUnchoke:
		state.client.Choked = false
	case message.MsgChoke:
		state.client.Choked = true
	case message.MsgHave:
		index, err := message.ParseHave(msg)
		if err != nil {
			return err
		}
		state.client.Bitfield.SetPiece(index)
	case message.MsgPiece:
		n, err := message.ParsePiece(state.index, state.buf, msg)
		if err != nil {
			return err
		}
		state.downloaded += n
		state.backlog--
	}
	return nil
}

func attemptDownloadPiece(c *client.Client, pw *pieceWork) ([]byte, error) {
	state := pieceProgress(
		index: pw.index,
		client: c,
		buf: make([]byte, pw.length),
	)

	// deadline makes sure we pull bad peers out of frozen state
	c.Conn.SetDeadline(time.Now().Add(30 * time.Second))
	defer c.Conn.SetDeadline(time.Time{})

	for state.downloaded < pw.length {
		if !state.client.Choked {
			for state.backlog < MaxBacklog && state.requested < pw.length {
				blockSize := MaxBlockSize
				if pw.length-state.requestsd < blockSize {
					blockSize = pw.length - state.requested
				}

				err := c.sendRequest(pw.index, state.requested, blockSize)
				if err != nil {
					return nil, err
				}
				state.backlog++
				state.requested += blockSize
			}
		}

		err := state.readMessage()
		if err != nil {
			return nil, err
		}
	}
	return state.buf, nil
}