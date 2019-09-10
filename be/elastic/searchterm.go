package elastic

import (
	"context"
	"github.com/cosmycx/wikied/model"
	elasticLib "github.com/olivere/elastic/v7"
	"log"
)


// GetForTerm infopost/s ...
func (c Client) SearchTerm(searchTerm string) ([]model.InfoPost, error) {

	ctx := context.Background()

	// Search with a term query
	termQuery := elasticLib.NewTermQuery("content", searchTerm)

	searchResult, err := c.elasticClient.Search().
		Index(IndexName).   // search in index "twitter"
		Query(termQuery).   // specify the query
		//Sort("created_at", true). // sort by "user" field, ascending
		From(0).Size(10).   // take documents 0-9
		//Pretty(true).       // pretty print request and response JSON
		Do(ctx)             // execute

		if err != nil {
			log.Printf("Error during execution GetAll : %v\n", err)
			return nil, err
		} // .if

	return c.convertAllResults(searchResult), nil
} // .GetForTerm infopost/s



