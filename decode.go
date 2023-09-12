package main

import (
	"fmt"
	"strconv"
	"math/rand"
	"time"
)

func decrypt(key string, altext string) {
	rand.Seed(int64(rand.Intn(1000)))
	fmt.Println("\n\n\nDecryption Algorithm ---->")
	number := altext
	input := key
	loopp := len(key)

	fmt.Println("key is: " + input)
	fmt.Println("len is: " + strconv.Itoa(loopp))
	
	sixDigits := ""
	everysix := ""
	allpt := ""

	for index := 0; index < len(number); index += 6 {
		for j := index; j < index+6; j++ {
			sixDigits += string(number[j])
		}

		everysix = sixDigits
		c := ""
		numLengthsix := len(everysix)

		for t := 0; t < numLengthsix; t++ {
			if everysix[0] == '0' && t == 0 {
				t++
			}
			c += string(everysix[t])
		}

		ept, _ := strconv.Atoi(c)
		rand.Seed(int64(ept)) // Seed the random number generator with 'ept'
		uh, _ := strconv.Atoi(key)
		pt := string(ept / uh)
		cc := string(pt)
		allpt += cc

		fmt.Println("before is: " + everysix + "  after: " + c + "  Plain text is: " + pt + "  the original text is: " + cc)
		sixDigits = ""
	}
}

func cdecrypt(key string, altext string) (string, time.Duration) {
	start := time.Now()
	rand.Seed(int64(rand.Intn(1000)))
	number := altext
	
	sixDigits := ""
	everysix := ""
	allpt := ""

	for index := 0; index < len(number); index += 6 {
		for j := index; j < index+6; j++ {
			sixDigits += string(number[j])
		}

		everysix = sixDigits
		c := ""
		numLengthsix := len(everysix)

		for t := 0; t < numLengthsix; t++ {
			if everysix[0] == '0' && t == 0 {
				t++
			}
			c += string(everysix[t])
		}

		ept, _ := strconv.Atoi(c)
		rand.Seed(int64(ept)) // Seed the random number generator with 'ept'
		uh, _ := strconv.Atoi(key)
		pt := string(ept / uh)
		cc := string(pt)
		allpt += cc

		sixDigits = ""
			
	}
	return allpt, time.Since(start)
}





