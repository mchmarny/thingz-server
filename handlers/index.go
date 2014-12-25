package handlers

import (
	"fmt"
	"net/http"

	"github.com/mchmarny/thingz-server/config"
)

func HandleGetIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Thingz API: "+config.API_VERSION)
}
