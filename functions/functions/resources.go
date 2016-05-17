package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/etc/hostname")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close() // HL

	bytes := make([]byte, 1024)
	_, err = f.Read(bytes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}
