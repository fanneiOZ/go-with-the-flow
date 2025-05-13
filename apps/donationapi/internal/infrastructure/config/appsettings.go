package config

import "infrastructure/pkg/external/omise"

type AppSettings struct {
	HttpPort       string `env:"http_port"`
	OmiseApiConfig *omise.ApiConfig
}
