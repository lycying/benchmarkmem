package benchmarkmem

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor/processorhelper"
	"go.uber.org/zap"
)

type spanProcessor struct {
	config Config
}

const (
	// The value of "type" key in configuration.
	typeStr = "benchmarkmem"
)

var processorCapabilities = consumer.Capabilities{MutatesData: true}
var logger = &zap.Logger{}

func createDefaultConfig() config.Processor {
	return &Config{}
}

func createTracesProcessor(
	_ context.Context,
	set component.ProcessorCreateSettings,
	cfg config.Processor,
	nextConsumer consumer.Traces,
) (component.TracesProcessor, error) {
	return processorhelper.NewTracesProcessor(
		cfg,
		nextConsumer,
		processTraces,
		processorhelper.WithCapabilities(processorCapabilities))
}

// NewFactory creates a factory for the routing processor.
func NewFactory() component.ProcessorFactory {
	return component.NewProcessorFactory(typeStr, createDefaultConfig,
		component.WithTracesProcessorAndStabilityLevel(createTracesProcessor, component.StabilityLevelStable))
}

func processTraces(_ context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	//rss := td.ResourceSpans()
	logger.Warn("", zap.Int("#spans", td.SpanCount()))
	return td, nil
}
