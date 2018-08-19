package bitset

var bitsPerInt uint = 64

type bitSet struct {
	size    uint
	storage []uint64
}

type BitSet interface {
	Clear(uint)
	GetStorage() *[]uint64
	MemorySize() uint
	Set(uint)
	Size() uint
	Test(uint) bool
}

func (bs *bitSet) Clear(i uint) {
	(*bs).storage[i/bitsPerInt] &^= 1 << (i % bitsPerInt)
}

func (bs *bitSet) GetStorage() *[]uint64 {
	return &((*bs).storage)
}

func (bs *bitSet) MemorySize() uint {
	return uint(len((*bs).storage)) * (bitsPerInt / 8)
}

func (bs *bitSet) Set(i uint) {
	(*bs).storage[i/bitsPerInt] |= 1 << (i % bitsPerInt)
}

func (bs *bitSet) Size() uint {
	return (*bs).size
}

func (bs *bitSet) Test(i uint) bool {
	return (((*bs).storage[i/bitsPerInt]>>(i%bitsPerInt))&1 == 1)
}

func New(size uint) *bitSet {
	bs := bitSet{size, make([]uint64, (size/bitsPerInt)+1)}
	return &bs
}
