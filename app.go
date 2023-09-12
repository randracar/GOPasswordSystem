package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
	"log"
	"strings"
	"io/ioutil"
)



type Account struct {
	Service string
	Login string
	PASS string
}

func split_string(s string) []string {
	res := strings.Split(s, "|")
	return res
}	

func search_account(key string){
	var acc Account

	textscanner := bufio.NewScanner(os.Stdin)
	accsfile := "data/accounts.txt"
	readFile, err := os.Open(accsfile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Input a string to search for: ")
	textscanner.Scan()

	s := textscanner.Text()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		if strings.Contains(fileScanner.Text(), s) {
			accinfo := split_string(fileScanner.Text())
			acc.Service = accinfo[0]
			acc.Login = accinfo[1]
			acc.PASS = accinfo[2]
			dp, t := cdecrypt(key, acc.PASS)
			fmt.Println("----------")
			fmt.Println("Account service: ", acc.Service)
			fmt.Println("Account Login: ", acc.Login)
			fmt.Println("Account Pass: ", dp)
			fmt.Println("Time to decrypt: ", t) 
			fmt.Println("----------")
		}
	}
	fmt.Println("Searched for all the accounts/logins with those parameters. If nothing is found, check the accounts.txt file.")
}

func add_account(key string){
	var acc Account
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("----------")
	fmt.Print("Input the name of the service to add: ")
	scanner.Scan()

	acc.Service = scanner.Text()

	fmt.Print("\nInput your login: ")
	scanner.Scan()

	acc.Login = scanner.Text()

	fmt.Print("\nInput your password (will be encrypted): ")
	scanner.Scan()


	var dur time.Duration
	acc.PASS, dur = ccrypt(key, scanner.Text())

	fmt.Println("Took ", dur, " to encrypt")

	finalstring := fmt.Sprintf("\n%s|%s|%s", acc.Service, acc.Login, acc.PASS)

	body, err := ioutil.ReadFile("data/accounts.txt")
	if err != nil {
		log.Fatal(err)
	}
	ecb := string(body)
 
	totalstring := fmt.Sprintf("%s%s", ecb, finalstring)

	err = ioutil.WriteFile("data/accounts.txt", []byte(totalstring), 0)
	if err != nil {
		log.Fatal(err)
	}
}

func process_input(c string, key string) bool {
	if c == "1" {
		// Add a new account
		add_account(key)
	} else if c == "2" {
		// search for a account
		search_account(key)
	} else {
		// return false, quit
		return false
	}
	return true
}

func app(key string){
	scanner := bufio.NewScanner(os.Stdin)
	running := true
	start := time.Now()

	for running == true {
		fmt.Println("----------")
		fmt.Println("What action do you want to perform?")
		fmt.Println("1 - Add a new account/text")
		fmt.Println("2 - View a account data")
		fmt.Println("3 - Quit")
		fmt.Println("----------")
		fmt.Print("Insert here your choice: ")
		scanner.Scan()
		user_choice := scanner.Text()
		running := process_input(user_choice, key)
		if running == false {
			break
		}
	}
	fmt.Println("Quitting the program...")
	fmt.Print("Used the app for ", time.Since(start))
}
