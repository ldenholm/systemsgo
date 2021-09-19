package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DbQuery struct {
	logger *log.Logger
}

func NewDbQuery(l *log.Logger) *DbQuery {
	return &DbQuery{l}
}

func listTablesQueryHandler() string {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-west-2"),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	svc := dynamodb.NewFromConfig(cfg)

	// Build the request with its input parameters
	resp, err := svc.ListTables(context.TODO(), &dynamodb.ListTablesInput{
		Limit: aws.Int32(2),
	})
	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
	}

	fmt.Println("Tables:")
	for _, tableName := range resp.TableNames {
		fmt.Println(tableName)
	}

	return strings.Join(resp.TableNames, ",")
}

func (query *DbQuery) GetTables(rw http.ResponseWriter, r *http.Request) {
	query.logger.Println("Handle GET Products")
	tables := listTablesQueryHandler()

	// currently needs a fix
	rw.Write([]byte(tables))
}
