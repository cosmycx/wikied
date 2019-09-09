package elastic

import (
	"context"
	"encoding/json"
	"github.com/cosmycx/wikied/model"
	elasticLib "github.com/olivere/elastic/v7"
	"log"
)

// GetInfoPostById ...
func (c Client)  GetInfoPostById(infoPostId string) (model.InfoPost, error) {

	ctx := context.Background()

	var infoPost model.InfoPost

	// Get tweet with specified ID
	get1, err := c.elasticClient.Get().
		Index(IndexName).
		//Type("tweet").
		Id(infoPostId).
		Do(ctx)
	if err != nil {
		log.Printf("Error during execution GetInfoPostById, err: %v", err.Error())
		return infoPost, err
	}// .if

	//if get1.Found {
	//	fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	//}

	result, err := c.convOneGetResult(get1)
	if err != nil {
		log.Printf("Error converting result : %v\n", err)
		return infoPost, err
	}
	return result, nil
}// .GetInfoPostById

// convert search result to infopost ...
func (c Client) convOneGetResult(searchResult *elasticLib.GetResult) (model.InfoPost, error) {

	var infoPostObj model.InfoPost

	err := json.Unmarshal(searchResult.Source, &infoPostObj)
	if err != nil {
		log.Printf("Can't deserialize 'infopost' object : %v\n", err)
		return infoPostObj, err
	}

	return infoPostObj, nil
} // .convOneGetResult