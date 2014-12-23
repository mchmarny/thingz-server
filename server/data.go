package server

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	flux "github.com/influxdb/influxdb/client"
	"github.com/mchmarny/thingz-server/types"
)

const (
	MAX_RECORDS_PER_QUERY = 1000
)

var db *DataService

func init() {

	c, err := parseDBConfig()
	if err != nil {
		log.Fatalf("Invalid connection string: %v", err)
		panic(err)
	}

	client, err := flux.NewClient(c)

	if err != nil {
		log.Fatalf("Error while creating InfluxDB client: %v", err)
		panic(err)
	}

	db = &DataService{
		Config:       c,
		Client:       client,
		ColumnFilter: strings.Split(Config.Filter, ","),
	}

}

type DataService struct {
	Config       *flux.ClientConfig
	Client       *flux.Client
	ColumnFilter []string
}

func getSourceFilters(src string, min int) (*types.ThingResponse, error) {

	//log.Printf("Querying filters for %s", src)
	q := fmt.Sprintf(
		"select * from /^%s.*/ where time > now() - %dm limit %d",
		src, min, MAX_RECORDS_PER_QUERY)

	result, err := db.Client.Query(q)
	if err != nil {
		log.Fatalf("Error on query [%s] - %v", q, err.Error())
		return nil, err
	}

	resp := &types.ThingResponse{
		Timestamp:  time.Now(),
		Dimensions: make([]*types.Dimension, 0),
	}

	for i, r := range result {

		log.Printf("Result[%d]%s", i, r.Name)
		d := &types.Dimension{Name: r.Name}

		for j, s := range r.Columns {
			log.Printf("Column[%d]%s", j, s)
			if !arrayContains(db.ColumnFilter, s) {

				f, err := getFilterRange(r.Name, s, min)
				if err != nil {
					log.Fatalf("Error on filter range query [%s] - %v", q, err.Error())
					return nil, err
				}

				d.Filters = append(d.Filters, &types.FilterCommand{
					Metric: s,
					Filter: f,
				})
			} // if contains
		} // for columns

		resp.Dimensions = append(resp.Dimensions, d)

	}

	return resp, nil

}

func getFilterRange(src, col string, min int) (*types.Range, error) {

	// log.Printf("Querying filter ranges for %s", src)
	q := fmt.Sprintf(
		"select PERCENTILE(\"%s\",%d), PERCENTILE(\"%s\",%d) from \"%s\" where time > now() - %dm",
		col, Config.MetricFilterBelow, col, Config.MetricFilterAbove, src, min)

	result, err := db.Client.Query(q)
	if err != nil {
		log.Fatalf("Error on query [%s] - %v", q, err.Error())
		return nil, err
	}

	if len(result) != 1 {
		log.Fatalf("Expected one result, go %d", len(result))
		return nil, errors.New("Invalid result")
	}

	resp := &types.Range{
		Below: result[0].Points[0][1],
		Above: result[0].Points[0][2],
	}

	return resp, nil

}

func arrayContains(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// parseDBConfig parses connStr string into an InfluxDB config
//    http://user:password@127.0.0.1:8086/dbname
//    udp://user:password@127.0.0.1:4444/dbname
func parseDBConfig() (*flux.ClientConfig, error) {

	u, err := url.Parse(Config.DB)
	if err != nil {
		return nil, err
	}

	c := &flux.ClientConfig{}

	c.IsUDP = (u.Scheme == "udp")
	c.Host = u.Host
	c.Username = u.User.Username()
	p, _ := u.User.Password()
	c.Password = p
	c.Database = strings.Replace(u.Path, "/", "", -1)

	return c, nil
}
