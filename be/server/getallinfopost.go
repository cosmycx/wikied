package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) getAllUser(w http.ResponseWriter, r *http.Request) {

	posts, err := s.ElasticClient.GetAll()
	if err != nil {
		log.Printf("Error retrieve all posts, err: %v\n", err)
		http.Error(w, "Error retrieve all posts", http.StatusInternalServerError)
		return
	} // .if

	json.NewEncoder(w).Encode(&posts)
} // .getAllUser
