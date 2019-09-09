package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) getInfoPostById(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not accepted", http.StatusMethodNotAllowed)
		return
	} // .if
	id := r.URL.Query()["id"]

	if len(id) == 0 {
		w.Write([]byte("empty payload"))
		w.WriteHeader(http.StatusBadRequest)
		return
	} // .if

	ids := id[0]

	// get infopost from Elastic
	infoPost, err := s.ElasticClient.GetInfoPostById(ids)
	if err != nil {
		log.Printf("Error retrieve infopost, err: %v\n", err)
		http.Error(w, "Error retrieve infopost", http.StatusNotFound)
		return
	} // .if

	json.NewEncoder(w).Encode(&infoPost)
} // .getAllUser
