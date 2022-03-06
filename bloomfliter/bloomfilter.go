package bloomfilter

const DefaultSize = 2 << 30

var (
	weis = [8]uint8{
		1 << 0,
		1 << 1,
		1 << 2,
		1 << 3,
		1 << 4,
		1 << 5,
		1 << 6,
		1 << 7,
	}
)

var seeds = []uint64{7, 11, 13, 31, 37}

type BloomFliter struct {
	set   *BitSet
	funcs [5]*SimpleHash
}

func NewBloomFilter() *BloomFliter {
	bf := new(BloomFliter)
	for i := 0; i < len(bf.funcs); i++ {
		bf.funcs[i] = &SimpleHash{DefaultSize, seeds[i]}
	}

	bf.set = NewBitSet(DefaultSize)
	return bf
}

func (b *BloomFliter) Add(v string) {
	for _, f := range b.funcs {
		b.set.Set(f.Hash(v))
	}
}

func (b *BloomFliter) Contains(v string) bool {
	result := true
	for _, f := range b.funcs {
		result = result && b.set.Contains(f.Hash(v))
	}
	return result
}

type SimpleHash struct {
	cap  uint64
	seed uint64
}

func (s *SimpleHash) Hash(str string) uint64 {
	var result uint64
	for i := 0; i < len(str); i++ {
		result = result*s.seed + uint64(str[i])
	}
	return (s.cap - 1) & result
}


type BitSet struct {
	cap uint64
	mem []uint8
}

func NewBitSet(cap uint64) *BitSet {
	byteNum := cap / 8
	var mem []uint8
	for i := uint64(0); i < byteNum; i++ {
		mem = append(mem, 0)
	}

	return &BitSet{
		cap: cap,
		mem: mem,
	}
}

func (b *BitSet) Set(num uint64) {
	byteIdx := num / 8
	weiIdx := num % 8
	b.mem[byteIdx] = b.mem[byteIdx] | weis[weiIdx]
}

func (b *BitSet) Contains(num uint64) bool {
	byteIdx := num / 8
	weiIdx := num % 8
	result := b.mem[byteIdx] & weis[weiIdx]
	return result == weis[weiIdx]
}
