package types

import (
	"time"
)

type Thingz struct {

	// Timestamp of when the metric was captured
	Timestamp time.Time `json:"ts"`

	// Sources pf metrics
	Sources []string `json:"thingz"`
}
