package server

import (
	"encoding/json"
	"github.com/cosmycx/wikied/model"
	"time"
	//	"io/ioutil"
	"log"
	"net/http"
)

func (s *Server) updateInfoPost(w http.ResponseWriter, r *http.Request) {

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

	if len(infoPost.Content) == 0 || len(infoPost.Id) == 0{
		w.Write([]byte("empty payload"))
		w.WriteHeader(http.StatusBadRequest)
		return
	} // .if

	// adding more info to infopost
	infoPost.CreatedAt = time.Now()

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
