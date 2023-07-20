package tracing

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

func Init(service string) (opentracing.Tracer, io.Closer) {
	//The tracer can be initialized with values coming from environment variables,
	//if it is built from a config that was created via FromEnv().
	//None of the env vars are required and all of them can be overridden via direct setting of the property on the configuration object.

	cfg, err := config.FromEnv()
	if err != nil {
		panic(fmt.Sprintf("ERROR: failed to read config from env vars: %v\n", err))
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}
