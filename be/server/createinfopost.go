package server

import (
	"context"
	"encoding/json"
	"github.com/cosmycx/wikied/elastic"
	"github.com/cosmycx/wikied/model"
	"github.com/google/uuid"
	"time"
	//	"io/ioutil"
	"log"
	"net/http"
)

func (s *Server) createInfoPost(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "method not accepted", http.StatusMethodNotAllowed)
		return
	} // .if

	var infoPost model.InfoPost

	err := json.NewDecoder(r.Body).Decode(&infoPost)
	if err != nil {
		log.Printf("Error converting post body to json, err: %v\n", err)
		http.Error(w, "payload json error", http.StatusBadRequest)
		return
	} // .if

	if len(infoPost.Content) == 0 || len(infoPost.Originator) == 0 {
		w.Write([]byte("empty payload"))
		w.WriteHeader(http.StatusBadRequest)
		return
	} // .if

	// adding more info to infopost
	infoPost.CreatedAt = time.Now()
	infoPost.Id = uuid.New().String()

	// inserting into elastic
	ctx := context.Background()

	err = s.ElasticClient.CreateIndexIfNotExists(ctx, elastic.IndexName)
	if err != nil {
		log.Printf("Error creating elastic index, err: %v\n", err)
		http.Error(w, "Error creating index", http.StatusInternalServerError)
		return
	} // .if

	// save this post in Elastic
	err = s.ElasticClient.InsertPost(infoPost)
	if err != nil {
		log.Printf("error inserting user/s, err: %v\n", err)
		http.Error(w, "error inserting user/s", http.StatusInternalServerError)
		return
	} // .if



	w.WriteHeader(http.StatusOK)
	type id struct {
		Id string `json:"id"`
	}// .id
	ids := id{
		Id: infoPost.Id,
	}// .ids
	json.NewEncoder(w).Encode(&ids)
	return
} // .add
