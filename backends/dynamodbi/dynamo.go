package dynamodbi

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/go-gorf/gorf"
)

type DynamoBackend struct {
	cfg aws.Config
}

// cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
//
//	if err != nil {
//		log.Fatalf("unable to load SDK config, %v", err)
//	}
func (b *DynamoBackend) Connect() (gorf.Db, error) {

	db := &DynamoDB{
		DB: dynamodb.NewFromConfig(b.cfg),
	}
	return db, nil
}

func (b *DynamoBackend) Disconnect() error {
	return nil
}
