package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// handle takes a json of two numbers, like this: {"x": 12, "y": 8}
// and returns a string with the sum, like this: sum=20
func handle(event Event) (string, error) {
	fmt.Println("version 3")
	sum := event.X + event.Y
	return fmt.Sprintf("sum=%d", sum), nil
}

func main() {
	lambda.Start(handle)
}
