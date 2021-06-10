package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
	"packSizeTest/service/packingList"
	"strconv"
)

func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("packingList handler invoked")
	// lets just imagine that packSizes was coming from a persistence 🥸
	packSizes := []int{250, 500, 1000, 2000, 5000}

	var requestedCount, err = strconv.Atoi(req.QueryStringParameters["requestedCount"])
	if err != nil {
		return clientError(http.StatusBadRequest)
	}
	log.Printf("packingList handler provided requestedCount:%d \n", requestedCount)

	packingListForRequest, err := packingList.GetPackingList(packSizes, requestedCount)
	if err != nil {
		log.Println("failed to generate packing list")
		// Not sure how to detect the "type" of error class with Golang here
		// so I will just blindly assume it is a client error 🤐
		return clientError(http.StatusBadRequest)
	} else {
		// I feel there is a better way to error handle than this 🤔
		log.Println("Packing List:")
		log.Println(packingListForRequest)
	}

	packingListJson, err := json.Marshal(packingListForRequest)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(packingListJson),
	}, nil
}

// Below helpers stolen from: https://www.alexedwards.net/blog/serverless-api-with-go-and-aws-lambda
func serverError(err error) (events.APIGatewayProxyResponse, error) {
	log.Println("Returning server error to caller 🔥")
	log.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	log.Println("returning client error to caller")
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}


func main() {
	log.Println("Running")
	lambda.Start(HandleRequest)
}
