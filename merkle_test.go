package merkle

import (
	"bytes"
	"crypto/sha256"
	"testing"
)

func getHasher() *Hasher {
	return NewHasher(func(b []byte) []byte {
		h := sha256.Sum256(b)
		return h[:]
	})
}

func TestEmptyTree(t *testing.T) {
	tree := NewMerkleTree(getHasher())
	if tree.LeafCount() != 0 {
		t.Fatal("Empty tree has nonzero leaf count")
	}
	if tree.LevelCount() != 0 {
		t.Fatal("Empty tree has nonzero level count")
	}
	if _, err := tree.LeafHash(0); err == nil {
		t.Fatal("Empty tree has nonnil leaf hash")
	}
}

func TestLeafHash(t *testing.T) {
	hasher := getHasher()
	tree := NewMerkleTree(hasher)

	leaf := []byte{0x1, 0x2, 0x3, 0x4}
	leafHash := hasher.HashLeaf(leaf)

	index := tree.AddLeaf(leaf)

	if hash, err := tree.LeafHash(index); err != nil {
		t.Fatal(err)
	} else if !bytes.Equal(hash, leafHash) {
		t.Fatal("Leaf hashes do not match: ", hash, leafHash)
	}

	if tree.LeafCount() != 1 {
		t.Fatal("Invalid leaf count")
	}
	if tree.LevelCount() != 1 {
		t.Fatal("Invalid level count")
	}

	tree.AddLeaf(leaf)
	if tree.LeafCount() != 2 {
		t.Fatal("Invalid leaf count")
	}
	if tree.LevelCount() != 2 {
		t.Fatal("Invalid level count")
	}

	tree.AddLeaf(leaf)
	if tree.LeafCount() != 3 {
		t.Fatal("Invalid leaf count")
	}
	if tree.LevelCount() != 3 {
		t.Fatal("Invalid level count")
	}

	tree.AddLeaf(leaf)
	if tree.LeafCount() != 4 {
		t.Fatal("Invalid leaf count")
	}
	if tree.LevelCount() != 3 {
		t.Fatal("Invalid level count")
	}
}
