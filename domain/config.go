package MetricsDomain

import (
	"strings"
)

type Config struct {
	Provider string `config:"METRICS_PROVIDER" default:"influxdb"`
}

func (c Config) IsDisabled() bool {
	return strings.ToLower(c.Provider) == "none"
}
