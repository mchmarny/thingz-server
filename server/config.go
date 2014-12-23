package server

import (
	"flag"
)

const (
	APP_VERSION = "0.0.1"
)

func init() {

	flag.IntVar(&Config.ServerPort, "port", 8080, "Server port")
	flag.StringVar(&Config.DBConnString, "db", "http://thingz:thingz@localhost:8086/thingz", "DB connection")
	flag.IntVar(&Config.DownsampleCCMin, "downsample", 5, "Minutes to downsample data for continuous queries")
	flag.IntVar(&Config.MetricFilterAbove, "metric-filter-above", 25, "Filter metrics above percentile")
	flag.IntVar(&Config.MetricFilterBelow, "metric-filter-below", 75, "Filter metrics below percentile")
	flag.IntVar(&Config.AgentCheckFreq, "agent-chech-freq", 60, "Agent check-in frequency in min")

	Config.Version = APP_VERSION

	flag.Parse()

}

// TODO: don't like this being public
var Config = &ServerConfig{}

type ServerConfig struct {
	Version           string
	ServerPort        int
	DBConnString      string
	DownsampleCCMin   int
	MetricFilterAbove int
	MetricFilterBelow int
	AgentCheckFreq    int
}
