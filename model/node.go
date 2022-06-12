package model

import (
	"github.com/anthoturc/baobabdb/hash"
	"net"
)

type Node struct {
	IpAddress net.IP
	ChordId   string
}

func NewNode(ip string) *Node {
	ipV4 := net.ParseIP(ip)
	return &Node{
		IpAddress: ipV4,
		ChordId:   hash.Hash(ip),
	}
}

