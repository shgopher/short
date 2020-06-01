package bloom

import "github.com/googege/godata"

type BloomFilter struct {
	bf *godata.BloomFilter
}

// return new bloom.
func NewBloom() *BloomFilter {
	return &BloomFilter{godata.NewBloomFilter(3)}
}
func (b *BloomFilter) Add(value []byte) {
	b.bf.Add(value)
}
func (b *BloomFilter) IsExit(value []byte) bool {
	return b.bf.IsExit(value)
}
