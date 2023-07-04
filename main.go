package main

import (
	"crypto/sha256"
	"fmt"
	"regexp"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Text string `json:"text"`
}

type MyResponse struct {
	Sha256 string `json:"sha256"`
}

const specialChars = "[`!@#$%^&*()_+-=\\[\\]{};':\"\\|,.<>\\/\\?~]"

var (
	pattern    = fmt.Sprintf("^(\\w|%s){8,}$", specialChars)
	oneNumber  = "\\d+"
	oneSpecial = fmt.Sprintf("%s+", specialChars)
)

func main() {

	lambda.Start(HandleLambdaEvent)

}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	isValid := validateString(event.Text)

	if !isValid {
		return MyResponse{Sha256: ""}, fmt.Errorf("the string is not valid")
	}

	sha256 := convertToSha256(event.Text)

	return MyResponse{Sha256: fmt.Sprintf("%x", sha256)}, nil
}

func validateString(s string) bool {

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

func convertToSha256(s string) []byte {
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	return bs
}
