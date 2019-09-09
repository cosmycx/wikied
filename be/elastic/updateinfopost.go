package elastic

import (
	"context"
	"github.com/cosmycx/wikied/model"
	"log"
)

// InsertPost/s ...
func (c Client) UpdateInfoPost(infoPost model.InfoPost) error {

	ctx := context.Background()

	// Index a infopost (using JSON serialization)
	_, err := c.elasticClient.Index().
		Index(IndexName).
		//Type("tweet").
		Id(infoPost.Id).
		BodyJson(infoPost).
		Do(ctx)
	if err != nil {
		log.Printf("Error inserting: %v\n", err)
		return err
	}// .err

	//log.Printf("Indexed infopost %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	// flush data need for refreshing data in index
	_, _ = c.elasticClient.Flush().Index(IndexName).Do(ctx)

	return err // nil
} // .InsertUsers
