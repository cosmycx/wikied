package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) searchTerm(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not accepted", http.StatusMethodNotAllowed)
		return
	} // .if

	searchTerms := r.URL.Query()["search"]
	log.Printf("searchTerms: %v\n", searchTerms)

	if len(searchTerms) == 0 {
		w.Write([]byte("missing search in query string"))
		w.WriteHeader(http.StatusBadRequest)
		return
	} // .if
	log.Printf("searchTerms[0]: %v\n", searchTerms[0])

	if len(searchTerms[0]) == 0 {
		w.Write([]byte("search is empty"))
		w.WriteHeader(http.StatusBadRequest)
		return
	} // .if

	// get infopost from Elastic
	infoPosts, err := s.ElasticClient.SearchTerm(searchTerms[0])
	if err != nil {
		log.Printf("Error retrieve infopost/s, err: %v\n", err)
		http.Error(w, "Error retrieve infopost/s", http.StatusNotFound)
		return
	} // .if

	json.NewEncoder(w).Encode(&infoPosts)
} // .getAllInfoPost
