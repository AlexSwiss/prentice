package controllers

import "github.com/AlexSwiss/prentice/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/shop/login", middlewares.SetMiddlewareJSON(s.ShopLogin)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/shops", middlewares.SetMiddlewareJSON(s.CreateShop)).Methods("POST")
	s.Router.HandleFunc("/shops", middlewares.SetMiddlewareJSON(s.GetShops)).Methods("GET")
	s.Router.HandleFunc("/shops/{id}", middlewares.SetMiddlewareJSON(s.GetShop)).Methods("GET")
	s.Router.HandleFunc("/shops/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateShop))).Methods("PUT")
	s.Router.HandleFunc("/shops/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteShop)).Methods("DELETE")
}
