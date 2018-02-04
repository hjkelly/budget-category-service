package main

import (
	"github.com/hjkelly/budget-category-service/common"
	"github.com/hjkelly/budget-category-service/views"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func main() {
	// Declare the secure routes. ---------
	router := httprouter.New()
	router.HandlerFunc("GET", "/v1/categories", views.ListCategories)
	router.HandlerFunc("POST", "/v1/categories", views.CreateCategory)

	// Attach middleware and serve!
	log.Fatal(http.ListenAndServe(":8080", wrapWithSecureMiddleware(router)))
}

// Take the finalized router and wrap it with the middleware we need.
func wrapWithSecureMiddleware(router *httprouter.Router) *negroni.Negroni {
	// Wrap the router with the above middlewares.
	m := negroni.New(
		// standard panic recovery middlware
		negroni.NewLogger(),
		negroni.NewRecovery(),
		common.GetAuth(),
	)
	m.UseHandler(router)
	return m
}
