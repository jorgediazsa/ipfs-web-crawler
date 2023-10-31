package api

import (
	"context"
	"net/http"
	"web-scraper/src/model"
)

type ContentAPIHandler interface {
	GetContentByIdHandler(w http.ResponseWriter, r *http.Request)
	GetContentsHandler(w http.ResponseWriter, r *http.Request)
	CreateContentHandler(w http.ResponseWriter, r *http.Request)
}

type ContentAPI interface {
	GetContentById(ctx context.Context, cid string) (model.Content, error)
	GetContents(ctx context.Context) ([]model.Content, error)
	CreateContent(ctx context.Context, data model.Content) error
}

type ContentDao interface {
	GetContentById(ctx context.Context, cid string) (model.Content, error)
	GetContents(ctx context.Context) ([]model.Content, error)
	SaveContent(ctx context.Context, data model.Content) error
}
