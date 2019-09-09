package elastic

import (
	"context"
	"errors"
	"fmt"
	elasticLib "github.com/olivere/elastic/v7"
)

// Ping method
func ping(ctx context.Context, client *elasticLib.Client, url string) error {

	// Ping the elasticsearch server to get HttpStatus, version number
	if client != nil {
		info, code, err := client.Ping(url).Do(ctx)
		if err != nil {
			return err
		}

		fmt.Printf("Elasticsearch returned with code %d and version %s \n", code, info.Version.Number)
		return nil
	}
	return errors.New("elastic client is nil")
}
