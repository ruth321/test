package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	h := sha1.New()
	_, _ = io.WriteString(h, "hello world")
	s := hex.EncodeToString(h.Sum(nil))[:5]
	fmt.Println(s)
}
