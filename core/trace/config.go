package trace

// TraceName represents the tracing name.
const TraceName = "go-tao"

// A Config is an opentelemetry config.
type Config struct {
	Name      string  `json:",optional"`
	AgentHost string  `json:",optional"`
	AgentPort string  `json:",optional"`
	Endpoint  string  `json:",optional"`
	Sampler   float64 `json:",default=1.0"`
	Batcher   string  `json:",default=jaeger,options=jaeger|zipkin|otlpgrpc|otlphttp"`
}
