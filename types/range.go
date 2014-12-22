package types

import (
	"fmt"
)

type Range struct {

	// Below range
	Below interface{} `json:"below"`

	// Above range
	Above interface{} `json:"above"`
}

func (r *Range) String() string {
	return fmt.Sprintf(
		"Range: [ Below:%v, Above:%v ]", r.Below, r.Above,
	)
}
