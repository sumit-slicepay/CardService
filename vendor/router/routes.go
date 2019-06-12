package router

import (
	controller "CardService/controller"
	"net/http"
)

// Route type description
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all routes
type Routes []Route

var routes = Routes{
	Route{
		"GetCVV",
		"GET",
		"/getCVV/{uuid}",
		*controller.CardController.GetCVV,
	},
}
