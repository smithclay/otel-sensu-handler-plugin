package main

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric/global"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	"time"

	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu/sensu-go/types"

	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpgrpc"

	"go.opentelemetry.io/otel/exporters/stdout"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
)

// Config represents the handler plugin config.
type Config struct {
	sensu.PluginConfig
}

var (
	plugin = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "otel-sensu-handler-plugin",
			Short:    "Generate OpenTelemetry metrics from Sensu",
			Keyspace: "sensu.io/plugins/otel-sensu-handler-plugin/config",
		},
	}
	options []*sensu.PluginConfigOption
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
	return fallback
	}
	return value
}

func main() {
	exporter, err := stdout.NewExporter(
		stdout.WithPrettyPrint(),
	)
	if err != nil {
		log.Fatalf("failed to initialize stdout export pipeline: %v", err)
	}

	ctx := context.Background()
	otelExporter, err := otlp.NewExporter(
		ctx,
		otlpgrpc.NewDriver(
			otlpgrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, "")),
			otlpgrpc.WithEndpoint(getenv("OTEL_EXPORTER_OTLP_METRIC_ENDPOINT", "ingest.lightstep.com:443")),
			otlpgrpc.WithHeaders(map[string]string{
				"lightstep-access-token":    os.Getenv("LS_ACCESS_TOKEN"),
			}),
		),
	)

	if err != nil {
		log.Fatalf("failed to initialize otelgrpc pipeline: %v", err)
	}

	pushController := controller.New(
		processor.New(
			simple.NewWithExactDistribution(),
			exporter,
		),
		controller.WithExporter(exporter),
		controller.WithExporter(otelExporter),
		controller.WithCollectPeriod(1*time.Second),
	)

	global.SetMeterProvider(pushController.MeterProvider())

	err = pushController.Start(ctx)
	if err != nil {
		log.Fatalf("failed to initialize metric controller: %v", err)
	}

	// Handle this error in a sensible manner where possible
	defer func() {
		log.Printf("stopping controller...")
		err = pushController.Stop(ctx)
		if err != nil {
			log.Fatalf("error stopping: %v", err)
		}
	}()

	log.Printf("starting sensu handler...")
	handler := sensu.NewGoHandler(&plugin.PluginConfig, options, checkArgs, executeHandler)
	handler.Execute()
}

func checkArgs(_ *types.Event) error {
	if len(os.Getenv("LS_ACCESS_TOKEN")) == 0 {
		return fmt.Errorf("LS_ACCESS_TOKEN is not set")
	}
	return nil
}

// based on: https://github.com/portertech/sensu-prometheus-pushgateway-handler/blob/main/main.go
func executeHandler(event *types.Event) error {
	meter := global.Meter("sensu-otel")
	for _, m := range event.Metrics.Points {
		var labels []attribute.KeyValue
		for _, t := range m.Tags {
			labels = append(labels, attribute.String(t.Name, t.Value))
		}
		recorder, err := meter.NewFloat64ValueRecorder(m.Name)
		if err != nil {
			return fmt.Errorf("error creating recorder: %v", err)
		}
		log.Printf("recording metric: %v=%v\n", m.Name, m.Value)
		recorder.Record(context.Background(), m.Value, labels...)
	}
	// HACK: Wait for metrics to flush
	time.Sleep(10 * time.Second)
	return nil
}
