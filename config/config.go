package config

import (
	"flag"
)

const (
	API_VERSION = "v1.0"
)

func init() {

	flag.IntVar(&Config.APIPort, "api-port", 8080, "API Server port")
	flag.BoolVar(&Config.Verbose, "verbose", false, "Output debug info")
	flag.IntVar(&Config.UIPort, "ui-port", 8081, "UI Server port")
	flag.StringVar(&Config.DBConnString, "db", "udp://thingz:thingz@localhost:4444/thingz", "DB connection")
	flag.IntVar(&Config.DownsampleCCMin, "downsample", 5, "Minutes to downsample data for continuous queries")
	flag.IntVar(&Config.MetricFilterAbove, "metric-filter-above", 25, "Filter metrics above percentile")
	flag.IntVar(&Config.MetricFilterBelow, "metric-filter-below", 75, "Filter metrics below percentile")
	flag.IntVar(&Config.AgentCheckFreq, "agent-chech-freq", 60, "Agent check-in frequency in min")
	flag.BoolVar(&Config.Load, "laod", false, "Whether server should load messages from Kafka")
	flag.StringVar(&Config.SubTopic, "topic", "thingz", "Kafka topic to subscribe to")
	flag.StringVar(&Config.SubBrokers, "brokers", "localhost:9092", "List of Kafka brokers and ports, use comma for multiple")
	flag.StringVar(&Config.PubDBConnString, "pub-db", "udp://thingz:thingz@localhost:4444/thingz", "Pub DB Connection")

	Config.Version = API_VERSION

	flag.Parse()

}

// TODO: don't like this being public
var Config = &ServerConfig{}

type ServerConfig struct {
	Version           string
	Verbose           bool
	APIPort           int
	UIPort            int
	DBConnString      string
	DownsampleCCMin   int
	MetricFilterAbove int
	MetricFilterBelow int
	AgentCheckFreq    int
	Load              bool
	SubTopic          string
	SubBrokers        string
	PubDBConnString   string
}
