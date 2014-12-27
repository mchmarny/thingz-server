package data

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/mchmarny/thingz-server/types"
)

// GetUtilizationByMetric
func GetUtilizationByMetric(group, metric string, min int) (*types.UtilizationResponse, error) {

	q := fmt.Sprintf(
		"select median(value) from /^src.*.dim.%s.met.%s/ where time > now() - %dm group by time(%dm) limit 1", group, metric, min, min)
	result, err := db.Query(q)
	if err != nil {
		log.Printf("Error on query [%s] - %v", q, err.Error())
		return nil, err
	}

	resp := &types.UtilizationResponse{
		Timestamp: time.Now().Unix(),
		Period:    makePeriod(min),
		Method:    "median",
	}

	list := types.ResourceUtilizationList{}

	for _, r := range result {
		list = append(list, types.ResourceUtilization{
			Resource: r.Name,
			Value:    r.Points[0][2],
		})
	}

	sort.Sort(list)

	resp.Resources = &list

	return resp, nil

}

func makePeriod(min int) *types.Period {
	p := &types.Period{
		From: time.Now().Unix(),
		To:   time.Now().Add(time.Duration(int32(min)) * time.Minute).Unix(),
	}
	return p
}
