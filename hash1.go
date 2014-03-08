package main

import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
)

func main() {
	mystr := "Hello World!"
	hash := md5.Sum([]byte(mystr))
	fmt.Println(hex.EncodeToString(hash[:]))
}
