package repository

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/ldenholm/systemsgo/data"
)

func initSession() *session.Session {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return sess
}

// dynamodb repository crud
func AddProduct(p *data.Product) (string, error) {
	sess := initSession()
	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Build attributes
	attVal, err := dynamodbattribute.MarshalMap(p)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}

	// Create item in table products
	//table := "products"
	input := &dynamodb.PutItemInput{
		Item:      attVal,
		TableName: aws.String("products"),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Printf("Got error calling PutItem: %s", err)
	}

	result := ("Successfully added '" + p.Name + "' to products")

	return result, err

}
