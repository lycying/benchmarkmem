package benchmarkmem

import (
	"go.opentelemetry.io/collector/config"
)

// Config represents the receiver config settings within the collector's config.yaml
type Config struct {
	config.ReceiverSettings `mapstructure:",squash"`
}
