package main

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

func generateLeaves(data []string) []*Node {
	leaves := make([]*Node, len(data))
	for i, d := range data {
		hash := DoubleSHA256([]byte(d))
		leaves[i] = &Node{hash: hash}
	}
	return leaves
}

func generateTree(leaves []*Node, hashFunc HashFunc) *Node {
	if len(leaves) == 1 {
		return leaves[0]
	}

	// handle odd number of leaves
	if len(leaves)%2 != 0 {
		leaves = append(leaves, leaves[len(leaves)-1])
	}

	var parents []*Node
	for i := 0; i < len(leaves); i += 2 {
		combinedHash := DoubleSHA256(append([]byte(leaves[i].hash), []byte(leaves[i+1].hash)...))
		parent := &Node{hash: combinedHash, left: leaves[i], right: leaves[i+1]}
		parents = append(parents, parent)
	}

	return generateTree(parents, hashFunc)
}

// func generateRecursiveTree(leaves []*Node, hash HashFunc) *Node {
// 	if len(leaves)%2 == 1 {
// 		return leaves[len(leaves)-1]
// 	}

// 	if len(leaves) == 2 {
// 		return NewNode(leaves[0].hash+leaves[1].hash, leaves[0], leaves[1])
// 	}

// 	mid := len(leaves) / 2
// 	left := generateRecursiveTree(leaves[:mid], hash)
// 	right := generateRecursiveTree(leaves[mid:], hash)
// 	generatedHash, _ := hash(left.hash, right.hash)

// 	return NewNode(generatedHash, left, right)
// }
