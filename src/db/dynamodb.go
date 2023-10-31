package db

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"go.uber.org/zap"
	"web-scraper/src/api"
	"web-scraper/src/model"
)

type Dao struct {
	Client api.AWSDynamoClient
	Logger *zap.Logger
}

func NewDao(client api.AWSDynamoClient, logger *zap.Logger) *Dao {
	return &Dao{
		Client: client,
		Logger: logger,
	}
}

func (d *Dao) SaveContent(ctx context.Context, data model.Content) error {
	item, err := attributevalue.MarshalMap(data)
	if err != nil {
		d.Logger.Error("error marshalling content", zap.Error(err))
		return err
	}

	_, err = d.Client.PutItem(ctx, &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("IPFSContent"),
	})
	if err != nil {
		d.Logger.Error("error putting content into dynamo", zap.Error(err))
		return err
	}

	return nil
}

func (d *Dao) GetContents(ctx context.Context) ([]model.Content, error) {
	params := &dynamodb.ScanInput{
		TableName: aws.String("IPFSContent"),
	}

	result, err := d.Client.Scan(ctx, params)
	if err != nil {
		d.Logger.Error("error scanning contents from dynamo", zap.Error(err))
		return nil, err
	}

	var contents []model.Content
	for _, item := range result.Items {
		var content model.Content
		if err := attributevalue.UnmarshalMap(item, &content); err != nil {
			d.Logger.Error("error unmarshalling content", zap.Error(err))
			return nil, err
		}
		contents = append(contents, content)
	}
	return contents, nil
}

func (d *Dao) GetContentById(ctx context.Context, cid string) (*model.Content, error) {
	params := &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"cid": &types.AttributeValueMemberS{Value: cid},
		},
		TableName: aws.String("IPFSContent"),
	}

	result, err := d.Client.GetItem(ctx, params)

	if err != nil {
		d.Logger.Error("error getting content", zap.Error(err))
		return nil, err
	}

	if len(result.Item) == 0 {
		return nil, nil
	}

	var content *model.Content
	if err := attributevalue.UnmarshalMap(result.Item, &content); err != nil {
		d.Logger.Error("error unmarshalling content", zap.Error(err))
		return nil, err
	}

	return content, nil
}
