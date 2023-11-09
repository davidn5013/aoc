package main

import (
	"crypto/md5"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	// fmt.Printf("%x\n", md5.Sum([]byte("abcdef609043")))
	fmt.Println(findmd5begin("bgvyzdsv", "00000"))
	fmt.Println(findmd5begin("bgvyzdsv", "000000"))
}

func findmd5begin(key, prefix string) (md5hash, answer string) {
	for i := 1; i < math.MaxInt; i++ {
		iAsStr := strconv.Itoa(i)
		md5 := fmt.Sprintf("%x", md5.Sum([]byte(key+iAsStr)))
		if strings.HasPrefix(md5, prefix) {
			return md5, iAsStr
		}
	}

	return "", ""
}
