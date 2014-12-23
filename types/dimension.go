package types

type Dimension struct {

	// Dimension of this source
	Dimension string `json:"dimension"`

	// Metric of this source
	Metric string `json:"metric"`

	// Filter filter
	Filter *Range `json:"filter"`
}
