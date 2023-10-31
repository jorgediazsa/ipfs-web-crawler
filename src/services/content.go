package services

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"web-scraper/src/db"
	"web-scraper/src/model"
)

type ContentService struct {
	dao    *db.Dao
	logger *zap.Logger
}

func NewContentService(dao *db.Dao, logger *zap.Logger) *ContentService {
	return &ContentService{
		dao:    dao,
		logger: logger,
	}
}

func (c *ContentService) GetContentById(ctx context.Context, cid string) (*model.Content, error) {
	if cid == "" {
		return nil, errors.New("invalid CID")
	}
	content, err := c.dao.GetContentById(ctx, cid)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (c *ContentService) GetContents(ctx context.Context) ([]model.Content, error) {
	contents, err := c.dao.GetContents(ctx)
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func (c *ContentService) CreateContent(ctx context.Context, data model.Content) error {
	err := c.dao.SaveContent(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
