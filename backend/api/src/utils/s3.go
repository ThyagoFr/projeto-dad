package utils

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/awserr"

	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	s3session *s3.S3
)

func createSession() {

	awsRegion := os.Getenv("AWS_REGION")

	s3session = s3.New(session.Must(
		session.NewSession(&aws.Config{
			Region: aws.String(awsRegion),
		}),
	))
	createBucket()

}

func createBucket() (err error) {

	awsBucket := os.Getenv("AWS_BUCKET_NAME")
	awsRegion := os.Getenv("AWS_REGION")

	_, err0 := s3session.CreateBucket(
		&s3.CreateBucketInput{
			Bucket: aws.String(awsBucket),
			CreateBucketConfiguration: &s3.CreateBucketConfiguration{
				LocationConstraint: aws.String(awsRegion),
			},
		},
	)

	if err0 != nil {
		if aerr, ok := err0.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				log.Println("Bucket name already exists!")
				err = err0
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				log.Println("Bucket name exists and is owned by you!")
			default:
				err = err0
			}
		}
	}
	return
}

// UploadToS3 - Upload image to S3
func UploadToS3(key, filename string) (string, error) {

	createSession()
	filename = "teste.png"
	f, err := os.Open(filename)
	if err != nil {
		log.Println("failed to open file")
		return "", err
	}

	awsBucket := os.Getenv("AWS_BUCKET_NAME")

	_, errF := s3session.PutObject(
		&s3.PutObjectInput{
			Body:   f,
			Bucket: aws.String(awsBucket),
			Key:    aws.String(key),
			ACL:    aws.String(s3.BucketCannedACLPublicRead),
		},
	)
	resource := "https://" + awsBucket + ".s3.amazonaws.com/" + key
	return resource, errF

}
