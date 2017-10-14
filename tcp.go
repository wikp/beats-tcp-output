package filebeat_tcp_output

import "github.com/elastic/beats/libbeat/outputs/codec"

type Config struct {
	Codec codec.Config `config:"codec"`

	BatchSize int
}

var defaultConfig = Config{}

