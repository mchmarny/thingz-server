package server

import (
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
	Config       *flux.ClientConfig
	Client       *flux.Client
	ColumnFilter []string
}

func getSourceFilters(src string) (*types.ThingResponse, error) {

	q := fmt.Sprintf("select low, high from /^%dm.%s.*/ limit 1", Config.DownsampleCCMin, src)
	result, err := db.Client.Query(q)
	if err != nil {
		log.Printf("Error on query [%s] - %v", q, err.Error())
		return nil, err
	}

	resp := &types.ThingResponse{
		Timestamp:  time.Now().Unix(),
		NextCheck:  time.Now().Add(time.Duration(int32(Config.AgentCheckFreq)) * time.Minute).Unix(),
		Dimensions: make([]*types.Dimension, 0),
	}

	if len(result) < 1 {

		ccErr := makeContinuousQuery(src)
		if ccErr != nil {
			log.Printf("Error during creation of continuous query %s", ccErr.Error())
			return nil, err
		}
		return resp, nil
	}

	for i, r := range result {
		log.Printf("Result[%d]%s", i, r.Name)
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

func makeContinuousQuery(src string) error {
	log.Printf("Creating continuous query for %s", src)
	q := fmt.Sprintf(
		"select min(value) as min, PERCENTILE(value, %d) as low, mean(value) as med, PERCENTILE(value, %d) as high, max(value) as max from /^%s.*/ group by time(%dm) into %dm.:series_name",
		Config.MetricFilterBelow, Config.MetricFilterAbove, src, Config.DownsampleCCMin, Config.DownsampleCCMin)
	_, err := db.Client.Query(q)
	return err
}

// parseDBConfig parses connStr string into an InfluxDB config
//    http://user:password@127.0.0.1:8086/dbname
//    udp://user:password@127.0.0.1:4444/dbname
func parseDBConfig() (*flux.ClientConfig, error) {

	u, err := url.Parse(Config.DBConnString)
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
