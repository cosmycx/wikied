package elastic

import (
	"context"
	elasticLib "github.com/olivere/elastic/v7"
	"log"
)

type Client struct {
	elasticClient *elasticLib.Client
} // .Client

// NewConn ..
func NewClient(ctx context.Context, url string, sniff bool) (*Client, error) {

	log.Printf("Initiating connection to elastic url: %v\n", url)
	//
	client, err := elasticLib.NewClient(elasticLib.SetURL(url), elasticLib.SetSniff(sniff))
	if err != nil {
		return nil, err
	} // .if

	err = ping(ctx, client, url)
	if err != nil {
		return nil, err
	} // .if

	ec := Client{
		elasticClient: client,
	} // .ec

	return &ec, nil
} // .NewClient
