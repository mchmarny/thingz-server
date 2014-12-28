package types

type ResourceUtilization struct {

	// Resource of when the metric was captured
	Resource string `json:"src"`

	// Value pf metrics
	Value interface{} `json:"val"`
}

type ResourceUtilizationList []ResourceUtilization

func (r ResourceUtilizationList) Len() int {
	return len(r)
}

func (r ResourceUtilizationList) Less(i, j int) bool {
	return r[i].Value.(float64) < r[j].Value.(float64)
}

func (r ResourceUtilizationList) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
