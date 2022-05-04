package main

import (
	"bytes"
	"encodings-ex01/base32"
	"encodings-ex01/base64"
	"fmt"
)

func main() {
	input := []byte("implement me")
	encoded64 := base64.Encode(input)
	decoded64 := base64.Decode(encoded64)
	fmt.Println(bytes.Equal(input, decoded64))
	fmt.Println(encoded64)
	fmt.Println(decoded64)
	encoded32 := base32.Base32encoder(input)
	decoded32 := []byte(base32.Base32decoder(input))
	fmt.Println(encoded32)
	fmt.Println(decoded32)
}
