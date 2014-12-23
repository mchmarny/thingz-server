package types

type Dimension struct {

	// Name of this source
	Name string `json:"name"`

	// Sources pf metrics
	Filters []*FilterCommand `json:"filters"`
}
