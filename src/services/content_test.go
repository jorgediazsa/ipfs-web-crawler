package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"go.uber.org/zap"
	"reflect"
	"testing"
	"web-scraper/src/db"
	"web-scraper/src/model"
)

func TestContentService_CreateContent(t *testing.T) {
	logger := zap.NewNop()
	type fields struct {
		dao    *db.Dao
		logger *zap.Logger
	}
	type args struct {
		data model.Content
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "save 123",
			fields: fields{
				dao: &db.Dao{
					Client: db.AWSDynamoClientMock{
						PutItemImpl: func(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
							return &dynamodb.PutItemOutput{}, nil
						},
					},
					Logger: logger,
				},
				logger: logger,
			},
			args: args{
				data: model.Content{
					CID: "123",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ContentService{
				dao:    tt.fields.dao,
				logger: tt.fields.logger,
			}
			if err := c.CreateContent(context.Background(), tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("CreateContent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContentService_GetContentById(t *testing.T) {
	logger := zap.NewNop()
	type fields struct {
		dao    *db.Dao
		logger *zap.Logger
	}
	type args struct {
		cid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Content
		wantErr bool
	}{
		{
			name: "get 123",
			fields: fields{
				dao: &db.Dao{
					Client: db.AWSDynamoClientMock{
						GetItemImpl: func(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
							return &dynamodb.GetItemOutput{
								Item: map[string]types.AttributeValue{
									"cid": &types.AttributeValueMemberS{Value: "123"},
								},
							}, nil
						},
					},
					Logger: logger,
				},
				logger: logger,
			},
			args: args{
				cid: "123",
			},
			want: &model.Content{
				CID: "123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ContentService{
				dao:    tt.fields.dao,
				logger: tt.fields.logger,
			}
			got, err := c.GetContentById(context.Background(), tt.args.cid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetContentById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetContentById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContentService_GetContents(t *testing.T) {
	logger := zap.NewNop()
	type fields struct {
		dao    *db.Dao
		logger *zap.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		want    []model.Content
		wantErr bool
	}{
		{
			name: "get all contents",
			fields: fields{
				dao: &db.Dao{
					Client: db.AWSDynamoClientMock{
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
					Logger: logger,
				},
				logger: logger,
			},
			want: []model.Content{
				{
					CID: "123",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ContentService{
				dao:    tt.fields.dao,
				logger: tt.fields.logger,
			}
			got, err := c.GetContents(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("GetContents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetContents() got = %v, want %v", got, tt.want)
			}
		})
	}
}
