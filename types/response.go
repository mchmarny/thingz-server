package types

import (
	"time"
)

type ThingResponse struct {

	// Timestamp of when the metric was captured
	Timestamp time.Time `json:"ts"`

	// Sources pf metrics
	Filters []FilterCommand `json:"filters"`
}
