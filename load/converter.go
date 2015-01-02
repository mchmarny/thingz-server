package load

import (
	"log"

	flux "github.com/influxdb/influxdb/client"
	"github.com/mchmarny/thingz-commons"
)

// convert translates the JSON data into a metric collection
// and then outputs each series individually
func convert(in <-chan []byte, out chan<- *flux.Series) {
	for {
		select {
		case data := <-in:
			m, err := commons.ParseMetric(data)
			if err != nil {
				log.Printf("Error parsing Metric Collection: %v", err)
			} else {
				out <- &flux.Series{
					Name:    m.FormatFQName(),
					Columns: []string{"time", "value"},
					Points:  [][]interface{}{{m.Timestamp.Unix(), m.Value}},
				} // series
			} // err
		} // select
	} // for
}
