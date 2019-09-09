package elastic

import (
	"context"
	"errors"
	"log"
)

// CreateIndexIfNotExists ...
func (c Client) CreateIndexIfNotExists(ctx context.Context, indexName string) error {

	exists, err := c.elasticClient.IndexExists(indexName).Do(ctx)
	if err != nil {
		log.Printf("Error checking index, err: %v\n", err)
		return err
	} // .if

	if exists {
		return nil
	} // .if

	res, err := c.elasticClient.CreateIndex(indexName).Do(ctx)

	if err != nil {
		return err
	} // .if

	if !res.Acknowledged {
		log.Printf("Error creating new index, err: %v\n", err)
		return errors.New("create index was not acknowledged, check that timeout value is correct")
	} // .if

	return nil
} // .CreateIndexIfNotExists
