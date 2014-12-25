package server

import (
	"net/http"

	"github.com/mchmarny/thingz-server/config"
	"github.com/mchmarny/thingz-server/handlers"
)

const (
	API_ROOT = "/api/" + config.API_VERSION
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", API_ROOT + "/", handlers.HandleGetIndex},
	Route{"GetFilter", "GET", API_ROOT + "/filters/{id}", handlers.HandleGetFilter},
}
