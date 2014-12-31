package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mchmarny/thingz-server/data"
)

// HandleGetUtilizationByMetric lists filters for this source
func HandleGetUtilizationByMetric(w http.ResponseWriter, r *http.Request) {

	SetResponseJSONEncoding(w)

	vars := mux.Vars(r)
	dimension := vars["dim"]
	metric := vars["met"]
	minute := vars["min"]

	if len(minute) < 1 || len(dimension) < 1 || len(metric) < 1 {
		WriteRequestError(w, "dimension, metric and min parameters required")
		return
	}

	m, err := strconv.Atoi(minute)
	if err != nil {
		WriteRequestError(w, "min parameter must be an integer")
		return
	}

	log.Printf("Thing utilization: %s.%s over %dm", dimension, metric, m)
	resp, err := data.GetUtilizationByMetric(dimension, metric, m)

	if err != nil {
		WriteRequestError(w, err.Error())
	} else {
		WriteResponse(w, resp)
	}
}
