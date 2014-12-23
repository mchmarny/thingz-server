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
		Config: c,
		Client: client,
	}

}

type DataService struct {
	Config *flux.ClientConfig
	Client *flux.Client
}

func query(q string, single bool) ([]*flux.Series, error) {

	result, err := db.Client.Query(q)
	if err != nil {
		log.Fatalf("Error on query [%s] - %v", q, err.Error())
		return nil, err
	}

	if single && len(result) != 1 {
		log.Fatalf("Expected single series, got [%d] from [%s]", len(result), q)
		return nil, errors.New("Invalid result")
	}

	log.Print(result)
	for i, r := range result {
		log.Printf("Result[%d]: %s", i, r.Name)
		log.Printf("   Columns: %v", r.Columns)
		log.Printf("   Points: %v", r.Points)
	}

	return result, nil

}

func getSourceFilters(src string, min int) (*types.ThingResponse, error) {

	log.Println("Querying filters for %s", src)
	q := fmt.Sprintf(
		"select * from /^%s.*/ where time > now() - %dm limit 1000",
		src, min)

	result, err := db.Client.Query(q)
	if err != nil {
		log.Fatalf("Error on query [%s] - %v", q, err.Error())
		return nil, err
	}

	resp := &types.ThingResponse{
		Timestamp: time.Now(),
	}
	resp.Dimensions = make([]*types.Dimension, len(result))

	for i, r := range result {

		log.Printf("Result[%d]%s", i, r.Name)

		d := &types.Dimension{
			Name: r.Name,
		}

		for j, s := range r.Columns {

			log.Printf("Column[%d]%s", j, s)

			d.Filters = append(d.Filters, &types.FilterCommand{
				Metric: s,
				Filter: &types.Range{},
			})
		}

		resp.Dimensions = append(resp.Dimensions, d)

	}

	return resp, nil

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
