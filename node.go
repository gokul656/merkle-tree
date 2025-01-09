package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

type HashFunc func(...string) (string, error)

// DoubleSHA256 performs Bitcoin-style double SHA-256 hashing
func DoubleSHA256(data []byte) string {
	hash1 := sha256.Sum256(data)
	hash2 := sha256.Sum256(hash1[:])
	return hex.EncodeToString(hash2[:])
}

func DefaultHashFunc(data ...string) (string, error) {
	b := strings.Builder{}
	for _, d := range data {
		bytes, _ := hex.DecodeString(d)
		hashString := DoubleSHA256(bytes)
		b.WriteString(hashString)
	}
	return b.String(), nil
}

type Node struct {
	hash  string // Stored in hex format
	left  *Node
	right *Node
}

func (n *Node) RootHash() string {
	bytes, _ := hex.DecodeString(n.hash)
	bytes = reverseBytes(bytes)
	return hex.EncodeToString(bytes)
}

func (n Node) Clone() *Node {
	return NewNode(n.hash, n.left, n.right)
}

func NewNode(hash string, left *Node, right *Node) *Node {
	return &Node{
		hash:  hash,
		left:  left,
		right: right,
	}
}
