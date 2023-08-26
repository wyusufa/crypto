package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

func DecodeHex(x string) string {
	log.Println("Start DecodeHex")

	src := []byte(x)                    // string to byte
	hexLength := hex.DecodedLen(len(x)) // panjang hex, setengah dari binary

	log.Println("decode length : ", hexLength)

	dst := make([]byte, hexLength)
	_, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("decode []byte : ", dst)
	log.Println("decode : ", string(dst))
	log.Println("End DecodeHex")

	return string(dst)
}

func DecodeBase64(x string) string {
	decodedByte, _ := base64.StdEncoding.DecodeString(x)
	decodedString = string(decodedByte)
	
	return ""
}

func IntToBinary(x int) {
	n := int64(x)

	fmt.Println(strconv.FormatInt(n, 2))
}

func BinaryToInt(x string) {
	if i, err := strconv.ParseInt(x, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
}

func main() {

}
