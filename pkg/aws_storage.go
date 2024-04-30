package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

type awsStorage struct {
	settings       *backend.DataSourceInstanceSettings
	customSettings dataSourceSettings
	query          dataSourceQuery
}

func newAWSStorage(_ context.Context, instance *dataSourceInstance, query dataSourceQuery, logger log.Logger) (*awsStorage, error) {
	logger.Error("Implement aws storage")

	customSettings, err := instance.Settings()
	if err != nil {
		return nil, err
	}
	logger.Error(fmt.Sprintf("customSettings: %s %s", customSettings.AWSAccesKeyId, customSettings.AWSSecretAccessKey))

	return &awsStorage{
		settings:       &instance.settings,
		customSettings: customSettings,
		query:          query,
	}, nil
}

// func lol(logger log.Logger, customSettings dataSourceSettings) {
// 	logger.Error("LISTING BUCKETS")

// 	sess := session.Must(session.NewSession(&aws.Config{
// 		Region:      aws.String("eu-central-1"),
// 		Credentials: credentials.NewStaticCredentials(customSettings.AWSAccesKeyId, customSettings.AWSSecretAccessKey, ""),
// 	}))

// 	// Create S3 service client
// 	svc := s3.New(sess)

// 	bucket := "testcsvvic"
// 	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String("testcsvvic")})
// 	if err != nil {
// 		logger.Error("Unable to list items in bucket %q, %v", bucket, err)
// 	}

// 	for _, item := range resp.Contents {
// 		logger.Error("Name:         ", *item.Key)
// 		logger.Error("Last modified:", *item.LastModified)
// 		logger.Error("Size:         ", *item.Size)
// 		logger.Error("Storage class:", *item.StorageClass)
// 		logger.Error("")
// 	}

// 	logger.Error("Found", len(resp.Contents), "items in bucket", bucket)
// 	logger.Error("")
// }

func (c *awsStorage) listFiles() (io.ReadCloser, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("eu-central-1"),
		Credentials: credentials.NewStaticCredentials(c.customSettings.AWSAccesKeyId, c.customSettings.AWSSecretAccessKey, ""),
	}))

	// Create S3 service client
	svc := s3.New(sess)

	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(c.customSettings.BucketName)})

	if err != nil {
		return nil, fmt.Errorf("Unable to list items in bucket %q, %v", c.customSettings.BucketName, err)
	}

	file, err := os.Create("test_list")

	if err != nil {
		return nil, fmt.Errorf("AWS: %s", err)
	}

	file.Write([]byte(fmt.Sprintf("filename,size\r\n")))
	for _, item := range resp.Contents {
		file.Write([]byte(fmt.Sprintf("%s,%d\r\n", *item.Key, *item.Size)))
	}

	file.Seek(0, io.SeekStart)

	return file, nil
}

func (c *awsStorage) Open() (io.ReadCloser, error) {
	if c.query.Experimental.ListDir {
		return c.listFiles()
	}

	if c.query.Path == "" {
		return nil, fmt.Errorf("AWS: need path")
	}

	file, err := os.Create(c.query.Path)
	if err != nil {
		return nil, fmt.Errorf("AWS: %s", err)
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("eu-central-1"),
		Credentials: credentials.NewStaticCredentials(c.customSettings.AWSAccesKeyId, c.customSettings.AWSSecretAccessKey, ""),
	}))

	downloader := s3manager.NewDownloader(sess)

	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(c.customSettings.BucketName),
			Key:    aws.String(c.query.Path),
		})

	return file, nil
}

func (c *awsStorage) Stat() error {
	return nil
}
