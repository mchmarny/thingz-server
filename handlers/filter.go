package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mchmarny/thingz-server/data"
	"github.com/mchmarny/thingz-server/types"
)

// HandleGetFilter lists filters for this source
func HandleGetFilter(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	log.Printf("Thing source: %s", id)
	resp, err := data.GetFilters(id)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(types.JSONError{
			Code: http.StatusBadRequest,
			Text: err.Error(),
		}); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
	}
}
