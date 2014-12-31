package load

import (
	flux "github.com/influxdb/influxdb/client"
	"github.com/mchmarny/thingz-server/config"
)

const (
	DATA_CHAN_NUM         = 5
	SERIES_CHAN_MULTPLIER = 20
)

// LoadFromKafka
func LoadFromKafka() {

	dataCh := make(chan []byte, DATA_CHAN_NUM)
	seriesCh := make(chan *flux.Series, DATA_CHAN_NUM)

	go func() {
		subscribe(
			config.Config.SubTopic,
			config.Config.SubBrokers,
			dataCh,
		)
	}()

	go func() {
		convert(
			dataCh,
			seriesCh,
		)
	}()

	go func() {
		publish(
			config.Config.PubDBConnString,
			seriesCh,
		)
	}()

}
