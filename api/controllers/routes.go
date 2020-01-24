package controllers

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/users", s.CreateUser).Methods("POST")
	s.Router.HandleFunc("/users", s.GetUsers).Methods("GET")
	s.Router.HandleFunc("/users/{id}", s.GetUser).Methods("GET")
	s.Router.HandleFunc("/users/{id}", s.UpdateUser).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", s.DeleteUser).Methods("DELETE")
	// s.Router.HandleFunc("/metrics", promhttp.Handler())
	// s.Router.HandleFunc("/metrics", promhttp.Handler())
	//Posts routes
	s.Router.HandleFunc("/posts", s.createPost).Methods("POST")
	s.Router.HandleFunc("/posts", s.GetPost).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", s.GetPost).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", s.UpdatePost).Methods("PUT")
	s.Router.HandleFunc("/posts/{id}", s.DeletePost).Methods("DELETE")
}
