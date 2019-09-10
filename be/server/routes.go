package server

import "net/http"

// RoutesInit ..
func (s *Server) RoutesInit() {

	s.Router.HandleFunc("/", s.home)

	s.Router.HandleFunc("/createinfopost", s.addAccessControl(s.createInfoPost))
	s.Router.HandleFunc("/updateinfopost", s.addAccessControl(s.updateInfoPost))

	s.Router.HandleFunc("/getallinfopost", s.addAccessControl(s.getAllInfoPost))

	s.Router.HandleFunc("/getinfopostbyid", s.addAccessControl(s.getInfoPostById))

} // .RoutesInit

// addAccessControl .. adds access control origins
func (s *Server) addAccessControl(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//TODO: dev only, to change
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == http.MethodOptions {
			//handle preflight
			w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			w.Header().Add("Access-Control-Max-Age", "86400")

			w.WriteHeader(http.StatusOK)
			return
		} // .if

		next.ServeHTTP(w, r) //proceed in the middleware chain

	}) // .http.HandlerFunc

} // .addAccessControl
