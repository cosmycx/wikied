package server

import (
	"fmt"
	"github.com/cosmycx/wikied/elastic"
	"net/http"
	"time"
)

type Server struct {
	// server router

	Router        *http.ServeMux
	ElasticClient *elastic.Client
} // .Server

// home |ping route, returns time
func (s *Server) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}// .if

	resp := fmt.Sprintf("server running time: %v", time.Now())

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
} // .home


