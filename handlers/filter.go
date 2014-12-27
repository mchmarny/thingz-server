package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mchmarny/thingz-server/data"
)

// HandleGetFilter lists filters for this source
func HandleGetFilter(w http.ResponseWriter, r *http.Request) {

	SetResponseJSONEncoding(w)

	vars := mux.Vars(r)
	src := vars["src"]

	if len(src) < 1 {
		WriteRequestError(w, "src parameter required")
		return
	}

	log.Printf("Thing source: %s", src)
	resp, err := data.GetFilters(src)

	if err != nil {
		WriteRequestError(w, err.Error())
	} else {
		WriteResponse(w, resp)
	}
}
