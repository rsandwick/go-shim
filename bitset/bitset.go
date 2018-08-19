package bitset

var bitsPerInt uint = 64

type bitSet struct {
	Size    uint
	Storage []uint64
}

func (bs *bitSet) Clear(i uint) {
	(*bs).Storage[i/bitsPerInt] &^= 1 << (i % bitsPerInt)
}

func (bs *bitSet) MemorySize() uint {
	return uint(len((*bs).Storage)) * (bitsPerInt / 8)
}

func (bs *bitSet) Set(i uint) {
	(*bs).Storage[i/bitsPerInt] |= 1 << (i % bitsPerInt)
}

func (bs *bitSet) Test(i uint) bool {
	return (((*bs).Storage[i/bitsPerInt]>>(i%bitsPerInt))&1 == 1)
}

func New(size uint) *bitSet {
	bs := bitSet{size, make([]uint64, (size/bitsPerInt)+1)}
	return &bs
}
