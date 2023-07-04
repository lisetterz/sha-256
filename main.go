package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/aws/aws-lambda-go/events"
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

	lambda.Start(HandleRequest)

}

func HandleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf(fmt.Sprintf("body:[%s] ", event.Body))

	var myEvent MyEvent

	err := json.Unmarshal([]byte(event.Body), &myEvent)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	isValid := validateString(myEvent.Text)

	if !isValid {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "",
		}, fmt.Errorf("the string %s from the event %s is not valid", myEvent.Text, myEvent)

	}

	sha256 := convertToSha256(myEvent.Text)

	msg := fmt.Sprintf("%x", sha256)

	responseBody := MyResponse{
		Sha256: msg,
	}

	jbytes, err := json.Marshal(responseBody)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	jstr := string(jbytes)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       jstr,
	}, nil

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
