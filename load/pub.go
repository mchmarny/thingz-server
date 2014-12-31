package load

import (
	"log"
	"net/url"
	"strings"

	flux "github.com/influxdb/influxdb/client"
)

// publish writes each series individually to avoid the
// UDP message size limitation
func publish(connStr string, in <-chan *flux.Series) {

	u, err := url.Parse(connStr)
	if err != nil {
		log.Fatalf("Error while parsing connection string: %v", connStr)
		panic(err)
	}

	c := &flux.ClientConfig{}
	c.IsUDP = (u.Scheme == "udp")
	c.Host = u.Host
	c.Username = u.User.Username()

	p, _ := u.User.Password()
	c.Password = p
	c.Database = strings.Replace(u.Path, "/", "", -1)

	client, err := flux.NewClient(c)
	if err != nil {
		log.Fatalf("Error while creating InfluxDB client: %v", err)
		panic(err)
	}

	var sendErr error

	for {
		select {
		case msg := <-in:
			if c.IsUDP {
				sendErr = client.WriteSeriesOverUDP([]*flux.Series{msg})
			} else {
				sendErr = client.WriteSeries([]*flux.Series{msg})
			} //
			if sendErr != nil {
				log.Printf("Error on send: %v", sendErr)
			}
		} // select
	} // for

}
