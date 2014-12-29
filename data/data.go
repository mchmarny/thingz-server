package data

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	flux "github.com/influxdb/influxdb/client"
	"github.com/mchmarny/thingz-server/config"
)

//    http://user:password@127.0.0.1:8086/dbname
//    udp://user:password@127.0.0.1:4444/dbname

func init() {

	u, err := url.Parse(config.Config.DBConnString)
	if err != nil {
		panic(err)
	}

	c := &flux.ClientConfig{}

	c.IsUDP = (u.Scheme == "udp")
	c.Host = u.Host
	c.Username = u.User.Username()
	p, _ := u.User.Password()
	c.Password = p
	c.Database = strings.Replace(u.Path, "/", "", -1)

	if len(config.Config.Proxy) > 1 {
		proxyUrl, err := url.Parse(config.Config.Proxy)
		if err != nil {
			log.Fatalf("Error while parsing HTTP proxy: %v", err)
			panic(err)
		}
		c.HttpClient = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	}

	client, err := flux.NewClient(c)
	if err != nil {
		log.Fatalf("Error while creating InfluxDB client: %v", err)
		panic(err)
	}

	db = client
}

var db *flux.Client
