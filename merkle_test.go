package merkle

import (
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
}
