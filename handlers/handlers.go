package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/miguelbelmar98/twittergo/models"
)

func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("Voy a procesar" + ctx.Value(models.Key("path")).(string) + ">" + ctx.Value(models.Key("method")).(string) + "<")

	var r models.RespApi
	r.Status = 400

	switch ctx.Value(models.Key("mathod")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		//TODO Por PATH
		}
		//TODO A la calma
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
		//TODO Por PATH
		}
		//TODO A la calma
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {
		//TODO Por PATH
		}
		//TODO A la calma
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {
		//TODO Por PATH
		}
		//TODO A la calma
	}
	r.Message = "Method Invalid"

	return r.CustomResponse, nil

}
