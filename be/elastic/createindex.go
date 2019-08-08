package elastic

import (
	"context"
	"errors"
)

// CreateIndexIfNotExists ...
func (c Client) CreateIndexIfNotExists(ctx context.Context, indexName string) error {

	exists, err := c.elasticClient.IndexExists(indexName).Do(ctx)
	if err != nil {
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
		return errors.New("CreateIndex was not acknowledged. Check that timeout value is correct.")
	} // .if

	return nil
} // .CreateIndexIfNotExists
