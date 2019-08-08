package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) getAllUser(w http.ResponseWriter, r *http.Request) {

	// get all users from Elastic
	ctx := context.Background()

	posts, err := s.ElasticClient.GetAll(ctx)
	if err != nil {
		log.Printf("Error retrieve all posts, err: %v\n", err)
		http.Error(w, "Error retrieve all posts", http.StatusInternalServerError)
		return
	} // .if

	json.NewEncoder(w).Encode(&posts)
} // .getAllUser
