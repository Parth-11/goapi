package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/Parth-11/goapi/internal/middleware"
)

func Handler(r *chi.Mux){
	//Global Middleware

	//Function strips the trailing slashes
	r.Use(chimiddle.StripSlashes)


	r.Route("/account",func(router chi.Router) {
		router.Use(middleware.Authorization)

		router.Get("/coins",GetCoinBalance)
	})
}