package awstools

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const AwsKeyEnv = "AWS_KEY"
const AwsSecretEnv = "AWS_SECRET"

type AwsConfigClient struct {
	config aws.Config
	region string
}

func NewAwsConfigClient(region string) (*AwsConfigClient, error) {
	key := os.Getenv(AwsKeyEnv)
	secret := os.Getenv(AwsSecretEnv)
	prov := credentials.NewStaticCredentialsProvider(key, secret, "")
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(prov),
	)
	if err != nil {
		return nil, err
	}
	return &AwsConfigClient{
		config: cfg,
		region: region,
	}, err
}

// DynamoDB builds a service client from the stored AWS config.
func (cli *AwsConfigClient) DynamoDB() *dynamodb.Client {
	if cli == nil {
		return nil
	}
	return dynamodb.NewFromConfig(cli.config)
}
