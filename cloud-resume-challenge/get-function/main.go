package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// var (
// 	// DefaultHTTPGetAddress Default Address
// 	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

// 	// ErrNoIP No IP found in response
// 	ErrNoIP = errors.New("No IP in HTTP response")

// 	// ErrNon200Response non 200 status code in response
// 	ErrNon200Response = errors.New("Non 200 Response found")
// )

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	var input = &dynamodb.GetItemInput{
		TableName: aws.String("cloud-resume-challenge"),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String("visitors"),
			},
		},
	}

	queryOutput, err := svc.GetItem(input)

	type Count struct {
		ID       string `json:"ID"`
		Visitors string `json:"visitors"`
	}
	count := Count{}

	err = dynamodbattribute.UnmarshalMap(queryOutput.Item, &count)

	if err != nil {
		log.Fatalf("Got error calling UpdateItem: %s", err)
	}

	// resp, err := http.Get(DefaultHTTPGetAddress)
	// if err != nil {
	// 	return events.APIGatewayProxyResponse{}, err
	// }

	// if resp.StatusCode != 200 {
	// 	return events.APIGatewayProxyResponse{}, ErrNon200Response
	// }

	// ip, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return events.APIGatewayProxyResponse{}, err
	// }

	// if len(ip) == 0 {
	// 	return events.APIGatewayProxyResponse{}, ErrNoIP
	// }

	return events.APIGatewayProxyResponse{
		// Body:       fmt.Sprintf("Hello, %v", string(ip)),
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "*",
			"Access-Control-Allow-Headers": "*",
		},
		Body:       fmt.Sprintf("{ \"count\": \"%s\" }", count.Visitors),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
