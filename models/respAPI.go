package models

import "github.com/aws/aws-lambda-go/events"

type RespApi struct {
	StatusCode     int
	Message        string
	CustomResponse *events.APIGatewayProxyResponse
}
