package elastic

import (
	"context"
	"encoding/json"
	"github.com/cosmycx/wikied/model"
	elasticLib "github.com/olivere/elastic/v7"
	"log"
)

// GetAll posts ...
func (c Client) GetAll() ([]model.InfoPost, error) {

	ctx := context.Background()

	query := elasticLib.MatchAllQuery{}

	searchResult, err := c.elasticClient.Search().
		Index(IndexName). // search in index
		Query(query).     // specify the query
		Do(ctx)           // execute
	if err != nil {
		log.Printf("Error during execution GetAll : %v\n", err)
		return nil, err
	} // .if

	return c.convertAllResults(searchResult), nil
} // .GetAll posts

// convert search result to infopost ...
func (c Client) convertAllResults(searchResult *elasticLib.SearchResult) []model.InfoPost {

	var result []model.InfoPost

	for _, hit := range searchResult.Hits.Hits {

			//log.Printf("hit.Source *&: %v\n", *&hit.Source)
			//log.Printf("hit.Source: %v\n", hit.Source)

		var infoPostObj model.InfoPost
		err := json.Unmarshal(hit.Source, &infoPostObj)
		if err != nil {
			log.Printf("Can't deserialize 'user' object : %s", err.Error())
			continue
		}
		result = append(result, infoPostObj)
	}
	return result
} // .convertAllResults