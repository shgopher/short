package bloom

import (
	"fmt"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	bf := NewBloom()
	bf.Add([]byte("1"))
	fmt.Println(bf.IsExit([]byte("1")))
	fmt.Println(bf.IsExit([]byte("12")))
}

func BenchmarkBloomFilter_Add(b *testing.B) {
	bf := NewBloom()
	for i := 0; i < b.N; i++ {
		bf.Add([]byte("1"))
	}
}

func BenchmarkBloomFilter_IsExit(b *testing.B) {
	bf := NewBloom()
	bf.Add([]byte("1"))
	for i := 0; i < b.N; i++ {
		bf.IsExit([]byte("1"))
	}
}