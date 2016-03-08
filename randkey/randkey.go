package randkey

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"sync/atomic"
)

var (
	base    string = randStr(4)
	counter uint64 = 0
)

// New return random key value build from two parts - process unique string and
// integer counter.
func New() string {
	n := atomic.AddUint64(&counter, 1)
	return fmt.Sprintf("%s:%d", base, n)
}

func randStr(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := base32.StdEncoding.EncodeToString(b)
	return s[:length]
}
