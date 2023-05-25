package short

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/shgopher/short/bloom"
	"github.com/shgopher/short/hash"
	qr "github.com/skip2/go-qrcode"
)

type Node struct {
	db         DB
	lock       sync.Mutex
	bf         *bloom.BloomFilter
	LongValue  string // long string
	ShortValue string // short string
}

func NewShort(db DB) *Node {
	return &Node{
		db: db,
		bf: bloom.NewBloom(),
	}
}

// short the long url to a small one in hash method.
func (n *Node) ShortAdd(longURL string) (string, error) {
	n.lock.Lock()
	defer n.lock.Unlock()
	murmurStringValue := hash.MurmurHash(longURL)
	n.LongValue = longURL
	n.ShortValue = murmurStringValue
	n.bf.Add([]byte(n.ShortValue))
	return n.db.Add(n)
}

// delete the shortURL,and it's long url
func (n *Node) ShortDelete(shortURL string) {
	n.lock.Lock()
	n.db.Delete(n.LongValue)
	n.lock.Unlock()
}

// find the longurl's shortURL.
func (n *Node) ShortFind(shortURL string) (string, error) {
	n.lock.Lock()
	defer n.lock.Unlock()
	_, short := filepath.Split(shortURL)
	// bloom filter
	if !n.bf.IsExit([]byte(short)) {
		return "", fmt.Errorf("bloom filter:no long url.")
	}
	return n.db.Find(short)
}

// change a new short URL.
func (n *Node) ShortChange(newShortURL string) error {
	n.lock.Lock()
	defer n.lock.Unlock()
	return n.db.Change(n, newShortURL)
}

// set shortURL to a QR CODE
func (n *Node) SetQR(url string, size int, file string) (err error) {
	n.lock.Lock()
	defer n.lock.Unlock()
	return qr.WriteFile(url+n.ShortValue, qr.Medium, size, file)
}
