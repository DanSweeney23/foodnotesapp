package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func CreateSessionToken(userId int) (res string) {
	now := time.Now()
	hashInput := fmt.Sprint(now.UnixNano()) + fmt.Sprint(userId)
	sessionToken := sha256.Sum256([]byte(hashInput))
	sessionId := hex.EncodeToString(sessionToken[:])

	return sessionId
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	userId := 1
	sessionId := CreateSessionToken(userId)

	return events.APIGatewayProxyResponse{
		Body:       sessionId,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
