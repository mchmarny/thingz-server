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

func ListThingz(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	log.Println("Getting sources...")
	getSources()
	resp := &types.Thingz{}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}

func GetThing(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	src := vars["src"]
	log.Printf("Thing src: %s", src)
	resp := &types.ThingResponse{}
	if resp != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
		return
	}

	// Can't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(types.JSONError{
		Code: http.StatusNotFound,
		Text: "Not Found",
	}); err != nil {
		panic(err)
	}

}
