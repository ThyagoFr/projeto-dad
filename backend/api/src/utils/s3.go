package utils

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws/awserr"

	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	s3session *s3.S3
)

func createS3Session(awsBucket string) {

	awsRegion := os.Getenv("AWS_REGION")

	s3session = s3.New(session.Must(
		session.NewSession(&aws.Config{
			Region: aws.String(awsRegion),
		}),
	))
	createBucket(awsBucket)

}

func createBucket(awsBucket string) (err error) {

	_, err0 := s3session.CreateBucket(
		&s3.CreateBucketInput{
			Bucket: aws.String(awsBucket),
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

// UploadBookCoverToS3 - Upload image to S3
func UploadBookCoverToS3(key, cover string) (string, error) {

	awsBucket := os.Getenv("AWS_BUCKET_NAME_BOOK")
	createS3Session(awsBucket)

	file, err := os.Create("temp.jpeg")

	response, err := http.Get(cover)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	file.Close()

	file, err = os.Open("temp.jpeg")

	_, errF := s3session.PutObject(
		&s3.PutObjectInput{
			Body:   file,
			Bucket: aws.String(awsBucket),
			Key:    aws.String(key),
			ACL:    aws.String(s3.BucketCannedACLPublicRead),
		},
	)
	if errF != nil {
		log.Fatal(errF)
	}
	resource := "https://" + awsBucket + ".s3.amazonaws.com/" + key
	return resource, errF

}

// UploadReaderProfileToS3 - UploadReaderProfileToS3
func UploadReaderProfileToS3(key uint, file multipart.File) (string, error) {

	awsBucket := os.Getenv("AWS_BUCKET_NAME_READER")
	createS3Session(awsBucket)

	_, errF := s3session.PutObject(
		&s3.PutObjectInput{
			Body:   file,
			Bucket: aws.String(awsBucket),
			Key:    aws.String(fmt.Sprint(key)),
			ACL:    aws.String(s3.BucketCannedACLPublicRead),
		},
	)
	if errF != nil {
		log.Fatal(errF)
	}
	resource := "https://" + awsBucket + ".s3.amazonaws.com/" + fmt.Sprint(key)
	return resource, errF

}
