package awstools

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Customer struct {
	ID       string `dynamodbav:"customer_id"`
	Email    string `dynamodbav:"email"`
	IsActive bool   `dynamodbav:"is_active"`
}

type AwsDynamoDbClient struct {
	ConfigCli *AwsConfigClient
	TableName string
}

func NewAwsDynamoDbClient(cfg *AwsConfigClient, tableName string) (*AwsDynamoDbClient, error) {
	return &AwsDynamoDbClient{
		ConfigCli: cfg,
		TableName: tableName,
	}, nil
}

func (cli *AwsDynamoDbClient) Save(item Customer) error {
	if cli == nil || cli.ConfigCli == nil {
		return errors.New("aws config client is nil")
	}
	if cli.TableName == "" {
		return errors.New("table name is required")
	}

	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return err
	}

	_, err = cli.ConfigCli.DynamoDB().PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: &cli.TableName,
		Item:      av,
	})
	return err
}

func (cli *AwsDynamoDbClient) Delete(id string) error {
	if cli == nil || cli.ConfigCli == nil {
		return errors.New("aws config client is nil")
	}
	if cli.TableName == "" {
		return errors.New("table name is required")
	}
	if id == "" {
		return errors.New("customer id is required")
	}

	_, err := cli.ConfigCli.DynamoDB().DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
		TableName: &cli.TableName,
		Key: map[string]types.AttributeValue{
			"customer_id": &types.AttributeValueMemberS{Value: id},
		},
	})
	return err
}

func (cli *AwsDynamoDbClient) Update(item Customer) error {
	if cli == nil || cli.ConfigCli == nil {
		return errors.New("aws config client is nil")

	}
	if item.ID == "" {
		return errors.New("customer id is required")
	}

	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return err
	}

	_, err = cli.ConfigCli.DynamoDB().PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: &cli.TableName,
		Item:      av,
	})
	return err
}

func (cli *AwsDynamoDbClient) Get(id string) (Customer, error) {
	if cli == nil || cli.ConfigCli == nil {
		return Customer{}, errors.New("aws config client is nil")
	}

	if id == "" {
		return Customer{}, errors.New("customer id is required")
	}

	result, err := cli.ConfigCli.DynamoDB().GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: &cli.TableName,
		Key: map[string]types.AttributeValue{
			"customer_id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return Customer{}, err
	}

	var customer Customer
	err = attributevalue.UnmarshalMap(result.Item, &customer)
	if err != nil {
		return Customer{}, err
	}

	return customer, nil
}
