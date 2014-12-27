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
	group := vars["group"]
	metric := vars["metric"]
	min := vars["min"]

	if len(min) < 1 || len(group) < 1 || len(metric) < 1 {
		WriteRequestError(w, "group, metric and min parameters required")
		return
	}

	m, err := strconv.Atoi(min)
	if err != nil {
		WriteRequestError(w, "min parameter must be an integer")
		return
	}

	log.Printf("Thing utilization: %s.%s over %dm", group, metric, m)
	resp, err := data.GetUtilizationByMetric(group, metric, m)

	if err != nil {
		WriteRequestError(w, err.Error())
	} else {
		WriteResponse(w, resp)
	}
}
