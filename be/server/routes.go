package server

// RoutesInit ..
func (s *Server) RoutesInit() {

	s.Router.HandleFunc("/", s.home)

	s.Router.HandleFunc("/createuser", s.createUser)

	s.Router.HandleFunc("/getalluser", s.getAllUser)

} // .RoutesInit
