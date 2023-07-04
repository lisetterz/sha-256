# sha-256

This API receives and input strings and validates it's format, it returns a sha 256 hash for a given string if and only if said string consists of at least 8 characters, at least one number, and at least one special character.

# Dependencies
Go 1.20

# Architecture
The RESTFul API code is written in Go and resides in an AWS Lambda function. It can be accessed through an API Gateway integration with Lambda.

![alt text](/data/flowDiagram.png)

# How to test

For more details you can find the OpenAPI [here](convertToSha256-API.yaml).

## Input format
The input formate is a json that follows the below schema like below.
```
{
    "text": "dfrff@34"
}
```

You should enter the string value that you want to convert into the text parameter.

Send a POST request with the json body to the endpoint below.
```
https://44gbi3ls84.execute-api.us-east-2.amazonaws.com/default/sha256Example
```

The CURL sample code is below:
```
curl --location 'https://44gbi3ls84.execute-api.us-east-2.amazonaws.com/default/sha256Example' \
--header 'Content-Type: application/json' \
--data-raw '{
    "text": "dfrff@34"
}'
```

## Output
The response is a json format containing the sha256 text parameter as the below sample:
```
{
    "sha256": "f50e0ea9347155396696c9772ede07dc12e074c2afef512f43b5c35310ce0e83"
}
```

In case that the string is invalid it will return an error format with a 404 Bad Request status code like below:
```
{
    "Message": "the string dfrff34 is not valid"
}
```

## Security
> **_NOTE:_**  For the purpose of this sample, the API is open for public to test. Later it will contain an API key that needs to be send in the an 'x-api-key' header.
No role is needed to test the application.