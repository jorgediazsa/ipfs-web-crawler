package api

import (
	"context"
	"web-scraper/src/model"
)

type ScraperAPI interface {
	Scrape(ctx context.Context, cid string) (model.Content, error)
}
