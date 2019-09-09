package server

import (
	"encoding/json"
	"github.com/cosmycx/wikied/model"
	"log"
	"net/http"
)

func (s *Server) getInfoPostById(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
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

	if len(infoPost.Id) == 0 {
		w.Write([]byte("empty payload"))
		w.WriteHeader(http.StatusBadRequest)
		return
	} // .if

	// get infopost from Elastic
	infoPost, err = s.ElasticClient.GetInfoPostById(infoPost.Id)
	if err != nil {
		log.Printf("Error retrieve infopost, err: %v\n", err)
		http.Error(w, "Error retrieve infopost", http.StatusNotFound)
		return
	} // .if

	json.NewEncoder(w).Encode(&infoPost)
} // .getAllUser
