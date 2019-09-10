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

	ids := r.URL.Query()["id"]
	if len(ids) == 0 {
		w.Write([]byte("missing id in query string"))
		w.WriteHeader(http.StatusBadRequest)
		return
	} // .if
	if len(ids[0]) == 0 {
		w.Write([]byte("id is empty"))
		w.WriteHeader(http.StatusBadRequest)
		return
	} // .if

	// get infopost from Elastic
	infoPost, err := s.ElasticClient.GetInfoPostById(ids[0])
	if err != nil {
		log.Printf("Error retrieve infopost, err: %v\n", err)
		http.Error(w, "Error retrieve infopost", http.StatusNotFound)
		return
	} // .if

	json.NewEncoder(w).Encode(&infoPost)
} // .getAllInfoPost
