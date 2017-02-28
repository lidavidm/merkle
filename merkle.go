package merkle

import "fmt"

type MerkleTree struct {
	hasher TreeHasher
	levels [][][]byte
}

func NewMerkleTree(hasher TreeHasher) *MerkleTree {
	return &MerkleTree{
		hasher: hasher,
		levels: nil,
	}
}

func (mt *MerkleTree) LeafCount() uint64 {
	if mt.levels != nil {
		return uint64(len(mt.levels[0]))
	} else {
		return 0
	}
}

func (mt *MerkleTree) LevelCount() uint64 {
	return uint64(len(mt.levels))
}

func (mt *MerkleTree) addLevel() {
	mt.levels = append(mt.levels, nil)
}

func (mt *MerkleTree) push(level int, node []byte) {
	mt.levels[level] = append(mt.levels[level], node)
}

func (mt *MerkleTree) AddLeaf(leaf []byte) uint64 {
	position := mt.LeafCount()

	// If we have reached the next power of two, add another level
	if isPowerOfTwoPlusOne(mt.LeafCount() + 1) {
		mt.addLevel()
	}

	hash := mt.hasher.HashLeaf(leaf)
	mt.push(0, hash)

	return position
}

type InvalidLeafError struct {
	index uint64
}

func (e InvalidLeafError) Error() string {
	return fmt.Sprintf("Leaf %d is not a valid leaf", e.index)
}

func (mt *MerkleTree) LeafHash(leaf uint64) ([]byte, error) {
	if leaf >= mt.LeafCount() {
		return nil, InvalidLeafError{leaf}
	}

	return mt.levels[0][int(leaf)], nil
}

func isPowerOfTwoPlusOne(n uint64) bool {
	if n == 0 {
		return false
	} else if n == 1 {
		return true
	} else {
		// Since n-1 would make it a power of 2, and (n-2) is
		// (n-1)-1, letting us reuse the usual bit trick
		return (n-1)&(n-2) == 0
	}
}

func CurrentRoot() ([]byte, error) {
	// TODO:
	return nil, nil
}
