# sha-256

This API receives and input strings and validates it's format, if the string is valid it returns the sha256 from the string.

# Dependencies
Go 1.20

# Architecture
The RESTFul API code is written in Go and resides in an AWS Lambda function. It can be accessed through an API Gateway integration with Lambda.



# How to test

For more details you can find the OpenAPI here.

## Input format
The input formate is a json that follows the below schema. You should enter the string value that you want to convert into the text parameter.

Send a POST request with the json body to the endpoint below.


The CURL sample code is below:

## Output
The response is a json format containing the sha256 text parameter.

## Security
For the purpose of this sample, the API is open for public to test. Later it will contain an API key that needs to be send in the an 'x-api-key' header.
No role is needed to test the application.