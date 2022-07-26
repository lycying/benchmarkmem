package benchmarkmem

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
)

const (
	// The value of "type" key in configuration.
	typeStr = "benchmarkmem"
)

func createDefaultConfig() config.Processor {
	return &Config{}
}

func createTracesProcessor(
	_ context.Context,
	set component.ProcessorCreateSettings,
	cfg config.Processor,
	nextConsumer consumer.Traces,
) (component.TracesProcessor, error) {
	return nil, nil
}

// NewFactory creates a factory for the routing processor.
func NewFactory() component.ProcessorFactory {
	return component.NewProcessorFactory(typeStr, createDefaultConfig,
		component.WithTracesProcessorAndStabilityLevel(createTracesProcessor, component.StabilityLevelStable))
}
