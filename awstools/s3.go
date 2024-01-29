package awstools

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"os"
	"path/filepath"
)

type AwsS3Client struct {
	config aws.Config
	bucket string
	s3     *s3.Client
}

func NewAwsS3Client(cfg *AwsConfigClient, bucket string) (*AwsS3Client, error) {
	client := s3.NewFromConfig(cfg.config)
	return &AwsS3Client{
		config: cfg.config,
		bucket: bucket,
		s3:     client,
	}, nil
}

func (cli *AwsS3Client) GetObjects() ([]types.Object, error) {
	output, err := cli.s3.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(cli.bucket),
	})
	if err != nil {
		return nil, err
	}
	data := make([]types.Object, len(output.Contents))
	for i, object := range output.Contents {
		data[i] = object
	}
	return data, nil
}

func (cli *AwsS3Client) UploadObject(path string) (*manager.UploadOutput, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	client := s3.NewFromConfig(cli.config)
	uploader := manager.NewUploader(client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(cli.bucket),
		Key:    aws.String(filepath.Base(path)),
		Body:   file,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (cli *AwsS3Client) Download(fileName string) (*manager.UploadOutput, error) {
	client := s3.NewFromConfig(cli.config)
	downloader := manager.NewDownloader(client)
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	_, err = downloader.Download(context.TODO(), file, &s3.GetObjectInput{
		Bucket: aws.String(cli.bucket),
		Key:    aws.String(fileName),
	})
	return nil, nil
}

func (cli *AwsS3Client) DeleteObject(fileName string) error {
	objectIds := []types.ObjectIdentifier{
		{Key: aws.String(fileName)},
	}
	_, err := cli.s3.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{
		Bucket: aws.String(cli.bucket),
		Delete: &types.Delete{Objects: objectIds},
	})
	return err
}
