package main
import ("fmt"
	"crypto/sha256"
	"encoding/hex")
func FindHash(diff uint, data string) (int, string) {
	for i := 0; true; i++ {
		out := sha256.Sum256([]byte(fmt.Sprintf("%s%d", data,i)))
		if checkBits(diff, out[:]) {return i,hex.EncodeToString(out[:])}
	}
	return 0,""
}
func main() { fmt.Println(FindHash(5, "Hello")) }
func checkBits(dif uint, b []byte) bool {
	for _,v := range b {
		if dif >= 8 {
			if v != 0 { return false }
			dif -= 8
		} else if dif > 0 {
			return (v & ^(0xff >> dif)) == 0
		}
	}
	return false
}
