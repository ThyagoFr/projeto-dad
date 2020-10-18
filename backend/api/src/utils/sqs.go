package utils

import (
	"errors"
	"os"

	"github.com/aws/aws-sdk-go/service/sqs"

	"github.com/aws/aws-sdk-go/aws/awserr"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	sqssession *sqs.SQS
)

func createSQSSession() (*string, error) {

	awsRegion := os.Getenv("AWS_REGION")

	sqssession = sqs.New(session.Must(
		session.NewSession(&aws.Config{
			Region: aws.String(awsRegion),
		}),
	))
	return createQueue()

}

func createQueue() (url *string, err error) {

	awsQueue := os.Getenv("AWS_QUEUE")

	_, err0 := sqssession.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String(awsQueue),
		Attributes: map[string]*string{
			"DelaySeconds":           aws.String("60"),
			"MessageRetentionPeriod": aws.String("86400"),
		},
	})

	if err0 != nil {
		if aerr, ok := err0.(awserr.Error); ok {
			switch aerr.Code() {
			case sqs.ErrCodeQueueNameExists:
				err = err0
			default:
				err = err0
			}
		}
	}
	result, _ := sqssession.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(awsQueue),
	})
	url = result.QueueUrl
	return
}

// SendMessage - Send message to a queue
func SendMessage(message Message) error {

	queueURL, _ := createSQSSession()
	_, err := sqssession.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"To": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(message.To),
			},
			"Name": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(message.Name),
			},
			"Token": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(message.Token),
			},
		},
		MessageBody: aws.String("Message struct with email request information"),
		QueueUrl:    queueURL,
	})
	if err != nil {
		return errors.New("Houve um erro ao enviar o email, tente novamente mais tarde")
	}
	return err

}
