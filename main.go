package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

//const specialChars = "/[`!@#$%^&*()_+-=[]{};':\"\\|,.<>/?~]/"

// var regexPattern = fmt.Sprintf("([0-9]+)([%s]+)[a-zA-Z0-9%s]{8,15}$", specialChars, specialChars)

var pattern = "^(\\w|[!@#$%^&*()]){8,}$"
var oneNumber = "\\d+"
var oneSpecial = "[!@#$%^&*()]+"

func main() {

	fmt.Println("Welcome to the Sha256 program.\nPlease enter the string to convert:")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\n", "", -1)

	valid := validateString(text)

	if valid {
		fmt.Println("The string is valid")
	} else {
		fmt.Println("The string is invalid")
	}

}

func validateString(s string) bool {

	if len(s) < 8 {
		return false
	}

	matched, err := regexp.MatchString(pattern, s)

	if err != nil {
		panic(err)
	}
	number, err := regexp.MatchString(oneNumber, s)

	if err != nil {
		panic(err)
	}

	special, err := regexp.MatchString(oneSpecial, s)

	if err != nil {
		panic(err)
	}

	if matched && number && special {
		return true
	}

	return false
}
