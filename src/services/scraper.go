package services

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"go.uber.org/zap"
	"web-scraper/src/model"
)

type ScraperService struct {
	logger         *zap.Logger
	IPFSGatewayURL string
}

func NewScraperService(logger *zap.Logger, IPFSGatewayURL string) *ScraperService {
	return &ScraperService{
		logger:         logger,
		IPFSGatewayURL: IPFSGatewayURL,
	}
}

func (s *ScraperService) Scrape(cid string) (model.Content, error) {
	var content model.Content
	var errorOnProcess error = nil

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		s.logger.Info("visiting", zap.String("URL", r.URL.String()))
	})

	c.OnError(func(_ *colly.Response, err error) {
		s.logger.Error("something went wrong", zap.Error(err))
		errorOnProcess = err
	})

	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &content)
		if err != nil {
			s.logger.Error("error marshalling content", zap.Error(err))
			errorOnProcess = err
		}
	})

	c.OnScraped(func(r *colly.Response) {
		s.logger.Info("finished scraping", zap.String("URL", r.Request.URL.String()))
	})

	url := fmt.Sprintf("%s/%s", s.IPFSGatewayURL, cid)
	err := c.Visit(url)
	if err != nil {
		return model.Content{}, err
	}

	if errorOnProcess != nil {
		return model.Content{}, errorOnProcess
	}

	return content, nil
}
