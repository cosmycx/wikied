package elastic

import (
	"context"
	"fmt"
	"github.com/cosmycx/wikied/model"
)

// InsertPost/s ...
func (c Client) InsertPost(infoPost model.InfoPost) error {

	ctx := context.Background()

	// Index a infopost (using JSON serialization)

	put1, err := c.elasticClient.Index().
		Index(IndexName).
		//Type("tweet").
		Id(infoPost.Id).
		BodyJson(infoPost).
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	// flush data need for refreshing data in index
	_, _ = c.elasticClient.Flush().Index(IndexName).Do(ctx)

	return err
} // .InsertUsers
