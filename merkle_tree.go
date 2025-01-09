package main

import "encoding/hex"

type MerkleTree struct {
	tree *Node
}

func (m *MerkleTree) RootHash() string {
	return m.tree.RootHash()
}

func (m *MerkleTree) PrintTree() {
	m.printRecTree(m.tree)
}

func (m *MerkleTree) printRecTree(node *Node) {
	if node == nil {
		return
	}

	m.printRecTree(node.left)
	m.printRecTree(node.right)
}

func NewMerkleTree(data []string, hashFunc HashFunc) *MerkleTree {
	leaves := generateLeaves(data)
	node := generateTree(leaves, hashFunc)
	return &MerkleTree{tree: node}
}

func NewDefaultMerkleTree(data []string) *MerkleTree {
	return NewMerkleTree(data, DefaultHashFunc)
}

func generateLeaves(txids []string) []*Node {
	leaves := make([]*Node, len(txids))
	for i, txid := range txids {
		// Convert hex string to bytes and reverse (Bitcoin's internal format)
		bytes, _ := hex.DecodeString(txid)
		bytes = reverseBytes(bytes)
		leaves[i] = &Node{hash: hex.EncodeToString(bytes)}
	}
	return leaves
}

func generateTree(leaves []*Node, hashFunc HashFunc) *Node {
	if len(leaves) == 1 {
		return leaves[0]
	}

	// handle odd number of leaves by duplicating last one
	if len(leaves)%2 != 0 {
		leaves = append(leaves, leaves[len(leaves)-1])
	}

	var parents []*Node
	for i := 0; i < len(leaves); i += 2 {
		leftBytes, _ := hex.DecodeString(leaves[i].hash)
		rightBytes, _ := hex.DecodeString(leaves[i+1].hash)

		// Concatenate the bytes and perform double SHA256
		combined := append(leftBytes, rightBytes...)
		combinedHash := DoubleSHA256(combined)

		parent := &Node{hash: combinedHash, left: leaves[i], right: leaves[i+1]}
		parents = append(parents, parent)
	}

	return generateTree(parents, hashFunc)
}

// Bitcoin uses little-endian byte order, so we need to reverse the bytes
func reverseBytes(b []byte) []byte {
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - 1 - i
		b[i], b[j] = b[j], b[i]
	}
	return b
}
