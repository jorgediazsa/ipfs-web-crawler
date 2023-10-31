## IPFS Web Scraper
IPFS Web Scraper written in go with REST api to handle requests
## Entry point
```
main.go
```
## Env vars
```
AWS_PROFILE //default: "personal"
AWS_REGION //default: "us-east-1"
HTTP_PORT //default: ":8000"
IPFS_GATEWAY_URL //default: "https://blockpartyplatform.mypinata.cloud/ipfs"
```
## How to run
```
go mod tidy
go run main.go
```
### then hit some of this endpoints
```
r.HandleFunc("/tokens", srv.GetContentsHandler)
r.HandleFunc("/tokens/{cid}", srv.GetContentByIdHandler)
r.HandleFunc("/scrape", srv.CreateContentHandler).Methods("POST")
r.HandleFunc("/bulk-scrape", srv.BulkCreateContentHandler).Methods("POST")
```
