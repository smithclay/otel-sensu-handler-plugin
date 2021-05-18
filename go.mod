module github.com/smithclay/otel-sensu-handler-plugin

go 1.14

require (
	github.com/sensu-community/sensu-plugin-sdk v0.11.0
	github.com/sensu/sensu-go/api/core/v2 v2.3.0
	github.com/sensu/sensu-go/types v0.3.0
	go.opentelemetry.io/otel v0.20.0
	go.opentelemetry.io/otel/exporters/otlp v0.20.0
	go.opentelemetry.io/otel/exporters/stdout v0.20.0
	go.opentelemetry.io/otel/metric v0.20.0
	go.opentelemetry.io/otel/sdk/metric v0.20.0
	google.golang.org/grpc v1.37.0
)
