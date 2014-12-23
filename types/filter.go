package types

import (
	"fmt"
)

type FilterCommand struct {

	// Dimension of this source
	Metric string `json:"metric"`

	// Filter filter
	Filter *Range `json:"filter"`
}

func (c *FilterCommand) String() string {
	return fmt.Sprintf(
		"FilterCommand: [ Metric:%s, Filter:%v ]", c.Metric, c.Filter,
	)
}
