package peers


import (
	"net"
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestUnmarshal(t *testing.T) {
	tests := map[string]struct {
		input string
		output []Peer
		fails bool
	}
	{
		"correctly parses peers": {
			input: string([]byte{127, 0, 0, 1, 0x00, 0x50, 1, 1, 1, 1, 0x01, 0xbb}),
			output: []Peer{
							{IP: net.IP{127, 0, 0, 1}, Port: 80},
							{IP: net.IP{1, 1, 1, 1}, Port: 443},
			}
		}
	}
}
