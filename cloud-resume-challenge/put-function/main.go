package main

import (
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("cloud-resume-challenge"),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String("visitors"),
			},
		},
		UpdateExpression: aws.String("ADD visitors :inc"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":inc": {
				N: aws.String("1"),
			},
		},
	}

	_, err := svc.UpdateItem(input)

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
		// Body:       fmt.Sprintf("{ \"count\": \"2\" }"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
