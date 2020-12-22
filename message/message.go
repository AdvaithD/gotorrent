package message

import (
	"encoding/binary"
	"fmt"
	"io"
)

type messageID uint8


const (
	// MsgChoke chokes the receiver
	MsgChoke messageID = 0
	//MsgUnchoke - unchoke the receiver
	MsgUnchoke messageID = 1
	// MsgInterested expresses interes in receiving data
	MsgInterested messageID = 2
	// MsgNotInterested expresses disinterest in receiving data
	MsgNoInterested messageID = 3
	// MsgHave - alerts the receiver that the sender has downloaded a piece
	MsgHave messageID = 4
	// MsgBitfield - encode which pieces  the sender has downloaded
	MsgBitfield messageID = 5
	// MsgRequest  - request s ablock of data from the receiver
	MsgRequest messageID = 6
	// MsgPiece - deliver a block of data that fulfills a request
	MsgPiece messageID = 7
	// MsgC ancel cancels a request
	MsgCancel messageID = 8
)

type Message struct {
	ID messageID
	Payload []byte
}

func FormatRequest(index, begin, legnth int) *Message {
	payload := make([]byte, 12)
}
