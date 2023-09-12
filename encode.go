package main

import (
	"fmt"
	"golang.org/x/crypto/sha3"
	"encoding/hex"
	"log"
    	"strconv"
    	"time"
)

func nothing(i int) {
	// so the compiler does not piss me off
}

func hash_string(s string) string {
	hash := sha3.New512()
	_, err := hash.Write([]byte(s))

	if err != nil {
		log.Fatal(err)
	}

	sha3 := hash.Sum(nil)

	es := hex.EncodeToString(sha3)
	return es
} 

// func to return char slice
func charslice(k string) []string {	
	slicerune := []rune(k)
	var char_slice []string
	for i, c := range slicerune {
		ec := string(c)
		char_slice = append(char_slice, ec)
		nothing(i)
	} 
	return char_slice
}

func crypt(key, plaintext string) string {

    fmt.Println("Encryption Algorithm ---->")
    fmt.Println("Plain Text: ", plaintext)

    randStr := key

	randNum, err := strconv.Atoi(key)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Key number is :", randStr)
    fmt.Println("Length of random number is :", len(randStr))

    var encryptedText string

    for i := 0; i < len(plaintext); i++ {
        character := plaintext[i]
        ascii := int(character)
        enctext := randNum * ascii
        tlenct := strconv.Itoa(enctext)
        lenct := len(tlenct)

        var s string
        if lenct < 6 {
            s = "0" + tlenct
            encryptedText += "0" + tlenct
        } else if lenct < 5 {
            s = "00" + tlenct
            encryptedText += "00" + tlenct
        } else if lenct < 4 {
            s = "000" + tlenct
            encryptedText += "000" + tlenct
        } else if lenct < 3 {
            s = "0000" + tlenct
            encryptedText += "0000" + tlenct
        } else if lenct < 2 {
            s = "00000" + tlenct
            encryptedText += "00000" + tlenct
        } else {
            encryptedText += tlenct
        }

        fmt.Printf("%c = %d     Encrypt Text is : %s\n", character, ascii, s)
    }

    altext := encryptedText
    fmt.Println("All cipher text is :", altext)
    return altext
}

func ToAscii(s string) string {
    var fs string
    for i, c := range s {
		if i >= 2 {
			ec := int(c)
			fs = fmt.Sprintf("%s%s", fs, strconv.Itoa(ec))
			break;
		} else if i < 2 {
			ec := int(c)
			fs = fmt.Sprintf("%s%s", fs, strconv.Itoa(ec))
		}
	}

    fs = fmt.Sprintf("%c%c%c", fs[0], fs[1], fs[2])
    return fs
}

func nothing2(s string){
    s = ""
    fmt.Print(s)
}

func ccrypt(key, plaintext string) (string, time.Duration) {
    start := time.Now()

    fmt.Println("Encryption Algorithm ---->")
    fmt.Println("Plain Text: ", plaintext)

    randStr := key

	randNum, err := strconv.Atoi(key)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Key number is :", randStr)
    fmt.Println("Length of random number is :", len(randStr))

    var encryptedText string

    for i := 0; i < len(plaintext); i++ {
        character := plaintext[i]
        ascii := int(character)
        enctext := randNum * ascii
        tlenct := strconv.Itoa(enctext)
        lenct := len(tlenct)

        var s string
        nothing2(s)
        if lenct < 6 {
            s = "0" + tlenct
            encryptedText += "0" + tlenct
        } else if lenct < 5 {
            s = "00" + tlenct
            encryptedText += "00" + tlenct
        } else if lenct < 4 {
            s = "000" + tlenct
            encryptedText += "000" + tlenct
        } else if lenct < 3 {
            s = "0000" + tlenct
            encryptedText += "0000" + tlenct
        } else if lenct < 2 {
            s = "00000" + tlenct
            encryptedText += "00000" + tlenct
        } else {
            encryptedText += tlenct
        }
    }

    altext := encryptedText
    fmt.Println("All cipher text is :", altext)
    return altext, time.Since(start)
}
