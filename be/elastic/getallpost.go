package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cosmycx/wikied/model"
	elasticLib "github.com/olivere/elastic/v7"
	"log"
)

// GetAll posts ...
func (c Client) GetAll(ctx context.Context) ([]model.Post, error) {

	query := elasticLib.MatchAllQuery{}

	searchResult, err := c.elasticClient.Search().
		Index(IndexName). // search in index
		Query(query).     // specify the query
		Do(ctx)           // execute
	if err != nil {
		fmt.Printf("Error during execution GetAll : %s", err.Error())
		return nil, err
	} // .if

	return convertResult(searchResult), nil
} // .GetAll posts

// convertSearchResultToPost ...
func convertResult(searchResult *elasticLib.SearchResult) []model.Post {

	var result []model.Post

	for _, hit := range searchResult.Hits.Hits {

		//	log.Printf("hit.Source *&: %v\n", *&hit.Source)
		//	log.Printf("hit.Source: %v\n", hit.Source)

		var postObj model.Post
		err := json.Unmarshal(*&hit.Source, &postObj)
		if err != nil {
			log.Printf("Can't deserialize 'user' object : %s", err.Error())
			continue
		}
		result = append(result, postObj)
	}
	return result
} // .convertResult
