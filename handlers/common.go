package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mchmarny/thingz-server/config"
	"github.com/mchmarny/thingz-server/types"
)

func HandleGetIndex(w http.ResponseWriter, r *http.Request) {
	SetResponseJSONEncoding(w)
	fmt.Fprint(w, "Thingz API: "+config.API_VERSION)
}

func SetResponseJSONEncoding(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func WriteRequestError(w http.ResponseWriter, e string) {
	w.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(w).Encode(types.JSONError{
		Code: http.StatusBadRequest,
		Text: e,
	}); err != nil {
		log.Fatalf("Error while encoding %s - %v", e, err.Error())
	}
}

func WriteResponse(w http.ResponseWriter, r interface{}) {
	w.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Fatalf("Error while encoding %s - %v", r, err.Error())
	}
}
