package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

type HashFunc func(...string) (string, error)

func DoubleSHA256(data []byte) string {
	hash1 := sha256.Sum256(data)
	hash2 := sha256.Sum256(hash1[:])
	return hex.EncodeToString(hash2[:])
}

func DefaultHashFunc(data ...string) (string, error) {
	b := strings.Builder{}
	for _, d := range data {
		hashString := DoubleSHA256([]byte(d))
		b.WriteString(hashString)
	}
	return b.String(), nil
}

type Node struct {
	hash  string
	left  *Node
	right *Node
}

func (n *Node) RootHash() string {
	return n.hash
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
