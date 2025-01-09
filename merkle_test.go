package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerkleGeneration(t *testing.T) {
	block100000 := []string{
		"8c14f0db3df150123e6f3dbbf30f8b955a8249b62ac1d1ff16284aefa3d06d87",
		"fff2525b8931402dd09222c50775608f75787bd2b87e56995a7bdd30f79702c4",
		"6359f0868171b1d194cbee1af2f16ea598ae8fad666d9b012c8ed2b79a236ec4",
		"e9a66845e05d5abc0ad04ec80f774a7e585c6e8db975962d069a522137b80c1d",
	}

	block100001 := []string{
		"bb28a1a5b3a02e7657a81c38355d56c6f05e80b9219432e3352ddcfc3cb6304c",
		"fbde5d03b027d2b9ba4cf5d4fecab9a99864df2637b25ea4cbcb1796ff6550ca",
		"8131ffb0a2c945ecaf9b9063e59558784f9c3a74741ce6ae2a18d0571dac15bb",
		"d6c7cb254aa7a5fd446e8b48c307890a2d4e426da8ad2e1191cc1d8bbe0677d7",
		"ce29e5407f5e4c9ad581c337a639f3041b24220d5aa60370d96a39335538810b",
		"45a38677e1be28bd38b51bc1a1c0280055375cdf54472e04c590a989ead82515",
		"c5abc61566dbb1c4bce5e1fda7b66bed22eb2130cea4b721690bc1488465abc9",
		"a71f74ab78b564004fffedb2357fb4059ddfc629cb29ceeb449fafbf272104ca",
		"fda204502a3345e08afd6af27377c052e77f1fefeaeb31bdd45f1e1237ca5470",
		"d3cd1ee6655097146bdae1c177eb251de92aed9045a0959edc6b91d7d8c1f158",
		"cb00f8a0573b18faa8c4f467b049f5d202bf1101d9ef2633bc611be70376a4b4",
		"05d07bb2de2bda1115409f99bf6b626d23ecb6bed810d8be263352988e4548cb",
	}

	mTree := NewDefaultMerkleTree(block100000)
	assert.Equal(t, mTree.RootHash(), "f3e94742aca4b5ef85488dc37c06c3282295ffec960994b2c0d5ac2a25a95766")

	mTree = NewDefaultMerkleTree(block100001)
	assert.Equal(t, mTree.RootHash(), "7fe79307aeb300d910d9c4bec5bacb4c7e114c7dfd6789e19f3a733debb3bb6a")
}
