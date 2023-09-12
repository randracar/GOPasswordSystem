package main

import (
	"fmt"
	"bufio"
	"os"
	"embed"
	"log"
)

//go:embed hash.txt
var hashfile embed.FS

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var uep string

	fmt.Println("----- PROOF OF CONCEPT - PASSWORD STORAGE APP -----")
	fmt.Printf("Input here your password: ")
	scanner.Scan()
	uep = scanner.Text()
	ep := hash_string(uep)

	text, err := hashfile.ReadFile("hash.txt")
	if err != nil {
		log.Fatal(err)
	}

	ectext := string(text)

	if ep == ectext {
		fmt.Println("The password is right!")
		uep := ToAscii(uep)
		app(uep)
	}
}
