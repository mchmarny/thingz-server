package data

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/mchmarny/thingz-server/config"
	"github.com/mchmarny/thingz-server/types"
)

func GetFilters(src string) (*types.ThingResponse, error) {

	q := fmt.Sprintf("select low, high from /^%dm.%s.*/ limit 1", config.Config.DownsampleCCMin, src)
	result, err := db.Query(q)
	if err != nil {
		log.Printf("Error on query [%s] - %v", q, err.Error())
		return nil, err
	}

	resp := &types.ThingResponse{
		Timestamp: time.Now().Unix(),
		NextCheck: time.Now().Add(
			time.Duration(int32(config.Config.AgentCheckFreq)) * time.Minute).Unix(),
		Dimensions: make([]*types.Dimension, 0),
	}

	if len(result) < 1 {

		ccErr := MakeContinuousFilterQuery(src)
		if ccErr != nil {
			log.Printf("Error during creation of continuous query %s", ccErr.Error())
			return nil, err
		}
		return resp, nil
	}

	for _, r := range result {
		//log.Printf("Result[%d]%s", i, r.Name)
		parts := strings.Split(r.Name, ".")
		d := &types.Dimension{
			Dimension: parts[2],
			Metric:    parts[3],
			Filter: &types.Range{
				Above: r.Points[0][2],
				Below: r.Points[0][3],
			},
		}
		resp.Dimensions = append(resp.Dimensions, d)
	}

	// TODO: add logic when paging
	resp.Count = len(resp.Dimensions)

	return resp, nil

}

func MakeContinuousFilterQuery(src string) error {
	log.Printf("Creating continuous query for %s", src)
	q := fmt.Sprintf(
		"select min(value) as min, PERCENTILE(value, %d) as low, mean(value) as med, PERCENTILE(value, %d) as high, max(value) as max from /^%s.*/ group by time(%dm) into %dm.:series_name",
		config.Config.MetricFilterBelow, config.Config.MetricFilterAbove, src, config.Config.DownsampleCCMin, config.Config.DownsampleCCMin)
	_, err := db.Query(q)
	return err
}
