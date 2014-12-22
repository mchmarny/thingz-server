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

	Config.Version = APP_VERSION

	flag.Parse()

}

// TODO: don't like this being public
var Config = &ServerConfig{}

type ServerConfig struct {
	Version string
	Port    int
	DB      string
}
