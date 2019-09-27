package storage

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/dominik-zeglen/inkster/config"
)

type AwsS3FileUploader struct {
	awsUploader *s3manager.Uploader
	bucket      string
}

func NewAwsS3FileUploader(config config.Config) AwsS3FileUploader {
	awsConfig := aws.Config{
		Region: &config.AWS.Region,
	}
	awsConfig.Credentials = credentials.NewStaticCredentials(
		config.AWS.AccessKey,
		config.AWS.SecretAccessKey,
		"",
	)

	sess, err := session.NewSession(&awsConfig)
	if err != nil {
		panic(err)
	}

	return AwsS3FileUploader{
		awsUploader: s3manager.NewUploader(sess),
		bucket:      config.Storage.S3Bucket,
	}
}

func (uploader AwsS3FileUploader) Upload(
	file io.Reader,
	filename string,
) (string, error) {
	input := s3manager.UploadInput{
		Bucket: &uploader.bucket,
		Key:    aws.String(filename),
		Body:   file,
		ACL:    aws.String("public-read"),
	}
	output, err := uploader.awsUploader.Upload(&input)

	if err != nil {
		return filename, err
	}

	return output.Location, nil
}
