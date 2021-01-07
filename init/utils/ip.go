package utils

import (
	"net"
)

// GetIP //
type GetIP interface {
	GetIP() (string, error)
}

// IP //
type IP struct {
}

// GetIP //
func (ip IP) GetIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	return addrs[1].String(), err
}
