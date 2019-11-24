package main

import "github.com/aws/aws-lambda-go/lambda"

type response struct {
  Message string `json:"message"`
}

func hello() (*response, error) {
  rs := &response{
    Message: "Hello world",
  }
	return rs, nil
}

func main() {
	lambda.Start(hello)
}
