package cfb8

import "crypto/cipher"

type cfb8 struct {
	b       cipher.Block
	next    []byte
	out     []byte
	pos     int

	decrypt bool
}

func (x *cfb8) XORKeyStream(dst, src []byte) {
	blockSize := x.b.BlockSize()
	for i := range src {
		blockEnd := x.pos + blockSize
		x.b.Encrypt(x.out, x.next[x.pos:blockEnd])

		if x.decrypt {
			x.next[blockEnd] = src[i]
		}
		dst[i] = src[i] ^ x.out[0]
		if !x.decrypt {
			x.next[blockEnd] = dst[i]
		}
		x.pos++

		if x.pos+blockSize == len(x.next) {
			copy(x.next, x.next[x.pos:])
			x.pos = 0
		}
	}
}

func NewCFB8Encrypter(block cipher.Block, iv []byte) cipher.Stream {
	return newCFB8(block, iv, false)
}

func NewCFB8Decrypter(block cipher.Block, iv []byte) cipher.Stream {
	return newCFB8(block, iv, true)
}

func newCFB8(block cipher.Block, iv []byte, decrypt bool) cipher.Stream {
	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		panic("cfb8.newCFB8: IV length must equal block size")
	}
	x := &cfb8{
		b:       block,
		next:    make([]byte, blockSize*4),
		out:     make([]byte, blockSize),
		pos:     0,
		decrypt: decrypt,
	}
	copy(x.next, iv)

	return x
}
