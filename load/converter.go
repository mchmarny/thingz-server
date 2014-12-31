package load

import (
	"fmt"
	"log"

	flux "github.com/influxdb/influxdb/client"
	types "github.com/mchmarny/thingz-commons"
)

// convert translates the JSON data into a metric collection
// and then outputs each series individually
func convert(in <-chan []byte, out chan<- *flux.Series) {
	for {
		select {
		case data := <-in:
			m, err := types.ParseMetricCollection(data)
			if err != nil {
				log.Printf("Error parsing Metric Collection: %v", err)
			} else {
				for _, v := range m.Metrics {
					out <- &flux.Series{
						Name: fmt.Sprintf("src.%s.dim.%s.met.%s",
							m.Source,
							m.Dimension,
							v.Metric,
						),
						Columns: []string{"time", "value"},
						Points:  [][]interface{}{{v.Timestamp.Unix(), v.Value}},
					} // series
				} // for
			} // err
		} // select
	} // for
}
