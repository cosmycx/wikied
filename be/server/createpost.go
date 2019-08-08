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

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "method not accepted", http.StatusMethodNotAllowed)
		return
	} // .if

	var postsIn []model.PostIn

	err := json.NewDecoder(r.Body).Decode(&postsIn)
	if err != nil {
		log.Printf("Error converting post body to json, err: %v\n", err)
		http.Error(w, "payload json error", http.StatusBadRequest)
		return
	} // .if

	if len(postsIn) == 0 {
		w.Write([]byte("no payload post/s"))
		w.WriteHeader(http.StatusBadRequest)
	} // .if

	// adding id and created_at
	posts := make([]model.Post, 0, len(postsIn))

	for _, pin := range postsIn {

		var post model.Post

		post.Originator = pin.Originator
		post.Content = pin.Content
		post.CreatedAt = time.Now()
		post.Id = uuid.New().String()
		if err != nil {
			log.Printf("error creating unique id, err: %v\n", err)
			http.Error(w, "server error creating id", http.StatusInternalServerError)
			return
		} // .if

		posts = append(posts, post)
	} // .for

	// inserting into elastic
	ctx := context.Background()

	err = s.ElasticClient.CreateIndexIfNotExists(ctx, elastic.IndexName)
	if err != nil {
		log.Printf("Error creating elastic index, err: %v\n", err)
		http.Error(w, "Error creating index", http.StatusInternalServerError)
		return
	} // .if

	// save this post in Elastic
	s.ElasticClient.InsertPost(ctx, posts)
	if err != nil {
		log.Printf("error inserting user/s, err: %v\n", err)
		http.Error(w, "error inserting user/s", http.StatusInternalServerError)
		return
	} // .if

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("inserted"))
	return
} // .add
