package builder

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"go.uber.org/zap"
	"service-worker-sqs-sns-postgres/core/domain"
	"service-worker-sqs-sns-postgres/dataproviders/awssqs"
	"service-worker-sqs-sns-postgres/dataproviders/consumer"
	repository "service-worker-sqs-sns-postgres/dataproviders/postgres/repository/events"
)

// NewSQS define all usecases to instantiate SQS.
func NewSQS(logger *zap.SugaredLogger, config *Configuration, session *session.Session, repo repository.IEventRepository) (domain.Source, error) {
	sqs, err := awssqs.NewSQSClient(session, config.SQSUrl, config.SQSMaxMessages, config.SQSVisibilityTimeout)
	if err != nil {
		return nil, fmt.Errorf("error awssqs.NewSQSClient: %w", err)
	}

	source, err := consumer.New(sqs, logger, config.SQSMaxMessages, repo)
	if err != nil {
		return nil, fmt.Errorf("error consumer.New: %w", err)
	}

	return source, nil
}
