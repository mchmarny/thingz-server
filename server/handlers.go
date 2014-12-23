package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mchmarny/thingz-server/types"
)

const (
	VERSION_TEXT = "v0.1.0"
	WELCOME_TEXT = "Thingz Server: " + VERSION_TEXT
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, WELCOME_TEXT)
}

// GetThing lists filters for this source
func GetThing(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	log.Printf("Thing source: %s", id)
	resp, err := getSourceFilters(id, Config.Scope)

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
