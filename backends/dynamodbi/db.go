package dynamodbi

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type DynamoDB struct {
	DB *dynamodb.Client
}

func (d *DynamoDB) Get(dest interface{}, key string) error {
	return nil
}

func (d *DynamoDB) Filter(dest interface{}, conds ...interface{}) error {
	return nil
}

func (d *DynamoDB) AutoMigrate(dst ...interface{}) error {
	return nil
}

func (d *DynamoDB) First(dest interface{}, conds ...interface{}) error {
	return nil
}
func (d *DynamoDB) Create(value interface{}) error {
	return nil
}
