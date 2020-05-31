package base62

import "testing"

var (
	intValue = 90999
)
func TestBase62(t *testing.T) {
	stringValue := EncodeBase62(intValue)
	curIntValue := DecodeBase62(stringValue)
	if curIntValue != intValue {
		t.Error("base62 encode and decode is error")
	}
}

func TestEncodeBase62(t *testing.T) {
	t.Log(EncodeBase62(100))
}

 func TestDecodeBase62(t *testing.T) {
	t.Log(DecodeBase62("a"))
}
//  29.6 ns/op
func BenchmarkEncodeBase62(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncodeBase62(100)
	}
}
//  80.3 ns/op
func BenchmarkEncodeBase622(b *testing.B) {
	for i:= 0;i< b.N;i++ {
		DecodeBase62("C1")
	}
}
