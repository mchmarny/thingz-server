package types

type ThingResponse struct {

	// Timestamp of when the metric was captured
	Timestamp int64 `json:"ts"`

	NextCheck int64 `json:"nextCheck"`

	// Count of total recrods
	Count int `json:"totalDimensions"`

	// HasMore indicates paging
	HasMore bool `json:"hasMore"`

	// Dimensions pf metrics
	Dimensions []*Dimension `json:"dimensions"`
}
