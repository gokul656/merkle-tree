# Bitcoin Merkle Tree Implementation

A Go implementation of Bitcoin's Merkle tree algorithm that correctly handles Bitcoin's internal representation and byte ordering requirements.

## Overview

This implementation provides:
- Correct byte ordering (internal little-endian, display big-endian)
- Bitcoin-style double SHA256 hashing
- Proper hash combinations for Merkle tree construction

## Usage

```go
txids := []string{
    "8c14f0db3df150123e6f3dbbf30f8b955a8249b62ac1d1ff16284aefa3d06d87",
    "fff2525b8931402dd09222c50775608f75787bd2b87e56995a7bdd30f79702c4",
    "6359f0868171b1d194cbee1af2f16ea598ae8fad666d9b012c8ed2b79a236ec4",
    "e9a66845e05d5abc0ad04ec80f774a7e585c6e8db975962d069a522137b80c1d",
}

// Create tree and get root hash
tree := NewDefaultMerkleTree(txids)
root := tree.RootHash()
```

## Implementation Notes

- Handles odd number of transactions by duplicating the last one
- Automatically manages Bitcoin's byte ordering requirements
- Matches Bitcoin Core's implementation for Merkle root calculation

## License

MIT License
