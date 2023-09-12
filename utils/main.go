package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("To Sha3-512 hasher")
	fmt.Print("Input here the string to hash: ")
	scanner.Scan()

	s := scanner.Text()

	hs := hash_string(s)

	fmt.Println(hs)
}
