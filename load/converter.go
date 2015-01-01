package load

import (
	"log"

	flux "github.com/influxdb/influxdb/client"
	"github.com/mchmarny/thingz-commons"
	"github.com/mchmarny/thingz-commons/types"
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
						Name:    commons.FormatMetricName(m.Source, m.Dimension, v.Metric),
						Columns: []string{"time", "value"},
						Points:  [][]interface{}{{v.Timestamp.Unix(), v.Value}},
					} // series
				} // for
			} // err
		} // select
	} // for
}
