package merkle

const (
	LeafPrefix = 0
	ChildPrefix = 1
)

type TreeHasher interface {
	HashEmpty() []byte
	HashLeaf(leaf []byte) []byte
	HashChildren(left, right []byte) []byte
}

type HasherFunc func([]byte) []byte

type Hasher struct {
	hasher HasherFunc
}

func NewHasher(hasher HasherFunc) *Hasher {
	return &Hasher {
		hasher: hasher,
	}
}

func (th Hasher) HashEmpty() []byte {
	return th.hasher([]byte{})
}

func (th Hasher) HashLeaf(leaf []byte) []byte {
	return th.hasher(append([]byte{LeafPrefix}, leaf...))
}

func (th Hasher) HashChildren(left, right []byte) []byte {
	return th.hasher(append(append([]byte{ChildPrefix}, left...), right...))
}
