package config

import (
	"flag"
)

const (
	API_VERSION = "v1.0"
)

func init() {

	flag.IntVar(&Config.APIPort, "api-port", 8080, "API Server port")
	flag.IntVar(&Config.UIPort, "ui-port", 8081, "UI Server port")
	flag.StringVar(&Config.DBConnString, "db", "http://thingz:thingz@localhost:8086/thingz", "DB connection")
	flag.IntVar(&Config.DownsampleCCMin, "downsample", 5, "Minutes to downsample data for continuous queries")
	flag.IntVar(&Config.MetricFilterAbove, "metric-filter-above", 25, "Filter metrics above percentile")
	flag.IntVar(&Config.MetricFilterBelow, "metric-filter-below", 75, "Filter metrics below percentile")
	flag.IntVar(&Config.AgentCheckFreq, "agent-chech-freq", 60, "Agent check-in frequency in min")
	flag.StringVar(&Config.Proxy, "proxy", "", "HTTP Proxy")

	Config.Version = API_VERSION

	flag.Parse()

}

// TODO: don't like this being public
var Config = &ServerConfig{}

type ServerConfig struct {
	Version           string
	APIPort           int
	UIPort            int
	DBConnString      string
	DownsampleCCMin   int
	MetricFilterAbove int
	MetricFilterBelow int
	AgentCheckFreq    int
	Proxy             string
}
