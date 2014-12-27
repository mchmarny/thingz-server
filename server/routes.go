package server

import (
	"github.com/mchmarny/thingz-server/config"
	"github.com/mchmarny/thingz-server/handlers"
	"github.com/mchmarny/thingz-server/types"
)

const (
	API_ROOT = "/api/" + config.API_VERSION
)

type Routes []types.Route

var routes = Routes{
	types.Route{"Index", "GET", API_ROOT + "/", handlers.HandleGetIndex},
	types.Route{"GetFilter", "GET", API_ROOT + "/filters/{src}", handlers.HandleGetFilter},
	types.Route{"GetSourcesByUtilization", "GET", API_ROOT + "/util/{group}/{metric}/{min}", handlers.HandleGetUtilizationByMetric},
}
