package types

import (
	"fmt"
)

type FilterCommand struct {

	// Dimension of this source
	Dimension string `json:"dimension"`

	// Filter filter
	Filter Range `json:"out"`
}

func (c *FilterCommand) String() string {
	return fmt.Sprintf(
		"FilterCommand: [ Dimension:%s, Filter:%v ]", c.Dimension, c.Filter,
	)
}
