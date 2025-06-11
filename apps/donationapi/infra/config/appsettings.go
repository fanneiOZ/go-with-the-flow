package config

import "sharedinfra/external/omise"

type AppSettings struct {
	HttpPort       string `env:"http_port"`
	OmiseApiConfig *omise.ApiConfig
}
