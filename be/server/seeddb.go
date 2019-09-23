package server

import (
	"time"

	"github.com/cosmycx/wikied/model"

	//	"io/ioutil"
	"log"
)

func (s *Server) SeedInfoPost(seedText string, seedId string) {

	var infoPost model.InfoPost

	// adding more info to infopost
	infoPost.CreatedAt = time.Now()
	infoPost.Id = seedId
	infoPost.Content = seedText

	// save this post in Elastic
	var err = s.ElasticClient.InsertPost(infoPost)
	if err != nil {
		log.Printf("error inserting user/s, err: %v\n", err)
		return
	} // .if

} // .add
