package awssns

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

// ClientSNS represents SNS client.
type ClientSNS struct {
	api sqsiface.SQSAPI
	url string
}

// NewSNSClient instances of a Client to connect SNS with session as parameter.
func NewSNSClient(sess *session.Session, url string) (*ClientSNS, error) {
	return &ClientSNS{
		api: sqs.New(sess),
		url: url,
	}, nil
}
