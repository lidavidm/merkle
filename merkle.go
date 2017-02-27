package merkle

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

// func (mt *MerkleTree) LeafHash(leaf uint64) ([]byte, error) {

// }
