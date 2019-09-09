package server

// RoutesInit ..
func (s *Server) RoutesInit() {

	s.Router.HandleFunc("/", s.home)

	s.Router.HandleFunc("/createinfopost", s.createInfoPost)

	s.Router.HandleFunc("/getallinfopost", s.getAllUser)

	s.Router.HandleFunc("/getinfopostbyid", s.getInfoPostById)

} // .RoutesInit
