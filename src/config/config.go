package config

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"go.uber.org/zap"
	"os"
	"web-scraper/src/db"
	server "web-scraper/src/rest"
	"web-scraper/src/services"
	"web-scraper/src/util"
)

type Config struct {
	DynamoDB       *dynamodb.Client
	Logger         *zap.Logger
	HttpPort       string
	IPFSGatewayURL string
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		panic("Config not set!")
	}
	return config
}

func InitializeServiceConfig(ctx context.Context) *server.Server {
	c := &Config{}

	awsRegion := util.StrDefault(os.Getenv("AWS_REGION"), "sa-east-1")
	awsProfile := util.StrDefault(os.Getenv("AWS_PROFILE"), "personal")
	c.DynamoDB = getDBConnection(ctx, awsRegion, awsProfile)

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	c.Logger = logger

	httpPort := util.StrDefault(os.Getenv("HTTP_PORT"), ":8000")
	c.HttpPort = httpPort

	IPFSGatewayURL := util.StrDefault(os.Getenv("IPFS_GATEWAY_URL"), "https://blockpartyplatform.mypinata.cloud/ipfs")
	c.IPFSGatewayURL = IPFSGatewayURL

	config = c

	dao := db.NewDao(c.DynamoDB, logger)

	return &server.Server{
		ContentService: services.NewContentService(dao, c.Logger),
		ScraperService: services.NewScraperService(logger, IPFSGatewayURL),
		Logger:         logger,
	}
}
