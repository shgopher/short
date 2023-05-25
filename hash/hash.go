package hash

import (
	"github.com/shgopher/short/base62"
	"github.com/spaolacci/murmur3"
)

//murmur hash.
func MurmurHash(longURL string) (shortURL string) {
	m := murmur3.Sum32([]byte(longURL))
	return base62.EncodeBase62(int(m))
}
