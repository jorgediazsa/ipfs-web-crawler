package db

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"go.uber.org/zap"
	"reflect"
	"testing"
	"web-scraper/src/model"
)

func Test_GetContentById(t *testing.T) {
	logger := zap.NewNop()

	type fields struct {
		client AWSDynamoClientMock
		logger *zap.Logger
	}
	type args struct {
		cid string
	}
	tests := []struct {
		name   string
		fields fields
		want   *model.Content
		args   args
	}{
		{
			name: "select 123",
			fields: fields{
				client: AWSDynamoClientMock{
					GetItemImpl: func(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
						return &dynamodb.GetItemOutput{
							Item: map[string]types.AttributeValue{
								"cid": &types.AttributeValueMemberS{Value: "123"},
							},
						}, nil
					},
				},
				logger: logger,
			},
			want: &model.Content{
				CID: "123",
			},
			args: args{cid: "123"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dao{
				Client: tt.fields.client,
				Logger: tt.fields.logger,
			}

			got, _ := d.GetContentById(context.Background(), tt.args.cid)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItem() got = %v, want %v", got, tt.want)
			}
		})
	}

}

func Test_GetContents(t *testing.T) {
	logger := zap.NewNop()

	type fields struct {
		client AWSDynamoClientMock
		logger *zap.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   []model.Content
	}{
		{
			name: "select all",
			fields: fields{
				client: AWSDynamoClientMock{
					ScanImpl: func(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {
						return &dynamodb.ScanOutput{
							Items: []map[string]types.AttributeValue{
								{
									"cid": &types.AttributeValueMemberS{Value: "123"},
								},
							},
						}, nil
					},
				},
				logger: logger,
			},
			want: []model.Content{
				{
					CID: "123",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dao{
				Client: tt.fields.client,
				Logger: tt.fields.logger,
			}

			got, _ := d.GetContents(context.Background())
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItem() got = %v, want %v", got, tt.want)
			}
		})
	}

}

func Test_SaveContent(t *testing.T) {
	logger := zap.NewNop()

	type fields struct {
		client AWSDynamoClientMock
		logger *zap.Logger
	}
	type args struct {
		data model.Content
	}
	tests := []struct {
		name   string
		fields fields
		want   error
		args   args
	}{
		{
			name: "select 123",
			fields: fields{
				client: AWSDynamoClientMock{
					PutItemImpl: func(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
						return &dynamodb.PutItemOutput{}, nil
					},
				},
				logger: logger,
			},
			want: nil,
			args: args{
				data: model.Content{CID: "123"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dao{
				Client: tt.fields.client,
				Logger: tt.fields.logger,
			}

			err := d.SaveContent(context.Background(), tt.args.data)
			if !reflect.DeepEqual(err, tt.want) {
				t.Errorf("GetItem() got = %v, want %v", err, tt.want)
			}
		})
	}

}
