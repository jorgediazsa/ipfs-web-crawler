package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"
	"web-scraper/src/config"
	server "web-scraper/src/rest"
)

var srv *server.Server

func main() {
	ctx := context.Background()
	srv = config.InitializeServiceConfig(ctx)
	r := mux.NewRouter()
	r.HandleFunc("/tokens", srv.GetContentsHandler)
	r.HandleFunc("/tokens/{cid}", srv.GetContentByIdHandler)
	r.HandleFunc("/scrape", srv.CreateContentHandler).Methods("POST")
	r.HandleFunc("/bulk-scrape", srv.BulkCreateContentHandler).Methods("POST")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		srv.Logger.Info(fmt.Sprintf("Listening on port %s", config.GetConfig().HttpPort))

		if err := http.ListenAndServe(config.GetConfig().HttpPort, r); err != nil {
			srv.Logger.Error("error running server", zap.Error(err))
		}
	}()

	<-stop

	srv.Logger.Info("Shutting down server...")
	context.WithTimeout(context.Background(), 30*time.Second)
	srv.Logger.Info("IPFS Content Scraper gracefully stopped")
}
