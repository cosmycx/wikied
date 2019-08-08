package elastic

import (
	"context"
	"fmt"

	"github.com/cosmycx/wikied/model"
)

// InsertPost/s ...
func (c Client) InsertPost(ctx context.Context, posts []model.Post) {

	for _, post := range posts {

		_, err := c.elasticClient.Index().Index(IndexName).Type(docType).BodyJson(post).Do(ctx)
		if err != nil {
			fmt.Printf("Post Id=%d was not created. Error : %s \n", post.Id, err.Error())
			continue
		} // .if

	} // .if

	// flush data need for refreshing data in index
	c.elasticClient.Flush().Index(IndexName).Do(ctx)
} // .InsertUsers
