package db

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type AWSDynamoClientMock struct {
	GetItemImpl        func(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	BatchGetItemImpl   func(ctx context.Context, params *dynamodb.BatchGetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.BatchGetItemOutput, error)
	BatchWriteItemImpl func(ctx context.Context, params *dynamodb.BatchWriteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.BatchWriteItemOutput, error)
	QueryImpl          func(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
	ScanImpl           func(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error)
	UpdateItemImpl     func(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)
	PutItemImpl        func(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	DeleteItemImpl     func(ctx context.Context, params *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error)
}

func (a AWSDynamoClientMock) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	if a.GetItemImpl != nil {
		return a.GetItemImpl(ctx, params, optFns...)
	}
	return &dynamodb.GetItemOutput{}, nil
}

func (a AWSDynamoClientMock) BatchGetItem(ctx context.Context, params *dynamodb.BatchGetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.BatchGetItemOutput, error) {
	if a.BatchGetItemImpl != nil {
		return a.BatchGetItemImpl(ctx, params, optFns...)
	}
	return &dynamodb.BatchGetItemOutput{}, nil
}

func (a AWSDynamoClientMock) BatchWriteItem(ctx context.Context, params *dynamodb.BatchWriteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.BatchWriteItemOutput, error) {
	if a.BatchWriteItemImpl != nil {
		return a.BatchWriteItemImpl(ctx, params, optFns...)
	}
	return &dynamodb.BatchWriteItemOutput{}, nil
}

func (a AWSDynamoClientMock) Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
	if a.QueryImpl != nil {
		return a.QueryImpl(ctx, params, optFns...)
	}
	return &dynamodb.QueryOutput{}, nil
}

func (a AWSDynamoClientMock) Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {
	if a.ScanImpl != nil {
		return a.ScanImpl(ctx, params, optFns...)
	}
	return &dynamodb.ScanOutput{}, nil
}

func (a AWSDynamoClientMock) UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error) {
	if a.UpdateItemImpl != nil {
		return a.UpdateItemImpl(ctx, params, optFns...)
	}
	return &dynamodb.UpdateItemOutput{}, nil
}

func (a AWSDynamoClientMock) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	if a.PutItemImpl != nil {
		return a.PutItemImpl(ctx, params, optFns...)
	}
	return &dynamodb.PutItemOutput{}, nil
}

func (a AWSDynamoClientMock) DeleteItem(ctx context.Context, params *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error) {
	if a.PutItemImpl != nil {
		return a.DeleteItemImpl(ctx, params, optFns...)
	}
	return &dynamodb.DeleteItemOutput{}, nil
}
