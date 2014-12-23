package server

import (
	"flag"
)

const (
	APP_VERSION = "0.0.1"
)

func init() {

	flag.IntVar(&Config.Port, "port", 8080, "Server port")
	flag.StringVar(&Config.DB, "db", "http://thingz:thingz@localhost:8086/thingz", "DB connection")
	flag.IntVar(&Config.Scope, "scope", 60, "Source filter scope in min")
	flag.StringVar(&Config.Filter, "filter", "time,sequence_number", "List of columns to filter out")
	flag.IntVar(&Config.MetricFilterAbove, "metric-filter-above", 20, "Filter metrics above percentile")
	flag.IntVar(&Config.MetricFilterBelow, "metric-filter-below", 80, "Filter metrics below percentile")

	Config.Version = APP_VERSION

	flag.Parse()

}

// TODO: don't like this being public
var Config = &ServerConfig{}

type ServerConfig struct {
	Version           string
	Port              int
	DB                string
	Scope             int
	Filter            string
	MetricFilterAbove int
	MetricFilterBelow int
}
