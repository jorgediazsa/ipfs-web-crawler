package config

import (
	"context"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
)

func getDBConnection(ctx context.Context, region string, profile string) *dynamodb.Client {
	cfg, err := awsConfig.LoadDefaultConfig(ctx, awsConfig.WithRegion(region), awsConfig.WithSharedConfigProfile(profile))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return dynamodb.NewFromConfig(cfg)
}
