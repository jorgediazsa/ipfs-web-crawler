package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"sync"
	"web-scraper/src/model"
	"web-scraper/src/services"
)

type Server struct {
	ContentService *services.ContentService
	ScraperService *services.ScraperService
	Logger         *zap.Logger
}

type ScraperRequestBody struct {
	CID string `json:"cid"`
}

type BulkScraperRequestBody struct {
	CIDs []string `json:"cids"`
}

type Content struct {
	Cid     string         `json:"cid"`
	Content *model.Content `json:"content"`
	Err     string         `json:"error"`
}

type ResponseBody []Content

func (s *Server) GetContentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	contents, err := s.ContentService.GetContents(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing request: %s", err), http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(contents)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshaling contents: %s", err), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(responseBody)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error writing http response: %s", err), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetContentByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid, ok := vars["cid"]
	if !ok {
		http.Error(w, "param cid is required", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	content, err := s.ContentService.GetContentById(r.Context(), cid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing request: %s", err), http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(&content)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshaling content: %s", err), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(responseBody)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error writing http response: %s", err), http.StatusInternalServerError)
		return
	}
}

func (s *Server) CreateContentHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody ScraperRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "param cid should be a string", http.StatusBadRequest)
		return
	}

	cid := requestBody.CID
	if cid == "" {
		http.Error(w, "param cid is required", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	content, err := s.ScraperService.Scrape(cid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing request: %s", err), http.StatusInternalServerError)
		return
	}

	content.CID = cid
	err = s.ContentService.CreateContent(r.Context(), content)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing request: %s", err), http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(content)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshaling content: %s", err), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(responseBody)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error writing http response: %s", err), http.StatusInternalServerError)
		return
	}

}

func (s *Server) BulkCreateContentHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody BulkScraperRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "param cids should be an array of strings", http.StatusBadRequest)
		return
	}

	cids := requestBody.CIDs
	cidsLen := len(cids)

	if cidsLen == 0 {
		http.Error(w, "param cids is required", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	wg := sync.WaitGroup{}

	ch := make(chan Content, cidsLen)
	for _, cid := range cids {
		wg.Add(1)
		go func(cidInner string) {
			defer wg.Done()

			content, err := s.ScraperService.Scrape(cidInner)
			if err != nil {
				s.Logger.Error("Error processing request", zap.Error(err))
				ch <- Content{Cid: cidInner, Content: nil, Err: err.Error()}
				return
			}

			content.CID = cidInner
			err = s.ContentService.CreateContent(r.Context(), content)
			if err != nil {
				s.Logger.Error("Error processing request", zap.Error(err))
				ch <- Content{Cid: cidInner, Content: &content, Err: err.Error()}
				return
			}
			ch <- Content{Cid: cidInner, Content: &content, Err: ""}
		}(cid)
	}

	s.Logger.Info(fmt.Sprintf("waiting for %d routines to finish", cidsLen))
	wg.Wait()
	close(ch)

	var rb ResponseBody
	for c := range ch {
		rb = append(rb, Content{
			Cid:     c.Cid,
			Content: c.Content,
			Err:     c.Err,
		})
	}

	responseBody, err := json.Marshal(rb)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshaling content: %s", err), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(responseBody)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error writing http response: %s", err), http.StatusInternalServerError)
		return
	}
}
