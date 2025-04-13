package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func main() {
	c := map[int]string{
		1: "0f4d0db3668dd58cabb9eb409657eaa8",
		2: "03eed983e18bb9ff9c6988fe4a7d41148fbb6594",
		3: "fb8e3f90a66bc874dfa0434cbbfe12c317ceff901c81c9b85c605229cc0dac82",
		4: "a9fa29c67637fdb8113e904fb8a58dc005afc753771be18e4453e07ac9fa634ace7acea679db57f645a2bc2b4a9729a6c49caf805a1534ec67ff7efbabd1e0b8",
	}

	fmt.Println("Welcome to the Login Portal.")
	fmt.Print("Please enter your password: ")

	var p string
	fmt.Scanln(&p)

	if len(p) < 16 {
		fmt.Println("Password too short.")
		return
	}

	a, b, d, e := p[:4], p[4:8], p[8:12], p[12:16]

	x := md5.Sum([]byte(a))
	y := sha1.Sum([]byte(b))
	z := sha256.Sum256([]byte(d))
	w := sha512.Sum512([]byte(e))

	if (hex.EncodeToString(x[:]) == c[1]) &&
		(hex.EncodeToString(y[:]) == c[2]) &&
		(hex.EncodeToString(z[:]) == c[3]) &&
		(hex.EncodeToString(w[:]) == c[4]) {
		fmt.Println("Access Granted.")
	} else {
		fmt.Println("Access Denied.")
	}
}
