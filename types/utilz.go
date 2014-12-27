package types

import (
	"log"
	"reflect"
)

type ResourceUtilization struct {

	// Resource of when the metric was captured
	Resource string `json:"src"`

	// Value pf metrics
	Value interface{} `json:"val"`
}

var floatType = reflect.TypeOf(float64(0))

type ResourceUtilizationList []ResourceUtilization

func (r ResourceUtilizationList) Len() int {
	return len(r)
}

func (r ResourceUtilizationList) Less(i, j int) bool {
	return getFloat(r[i].Value) < getFloat(r[j].Value)
}

func (r ResourceUtilizationList) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func getFloat(unk interface{}) float64 {
	v := reflect.ValueOf(unk)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(floatType) {
		log.Panicf("cannot convert %v to float64", v.Type())
		return 0
	}
	fv := v.Convert(floatType)
	return fv.Float()
}
