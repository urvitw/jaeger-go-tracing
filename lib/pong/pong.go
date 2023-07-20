package pong

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"net/http"
	libhttp "ping/lib/http"
	"ping/lib/tracing"
)

// Ping sends a ping request to the given hostPort, ensuring a new span is created
// for the downstream call, and associating the span to the parent span, if available
// in the provided context.

func Pong(ctx context.Context, hostPort string) (string, error) {

	span, _ := opentracing.StartSpanFromContext(ctx, "pong-send")
	defer span.Finish()

	url := fmt.Sprintf("http://%s/pong", hostPort)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	if err := tracing.Inject(span, req); err != nil {
		return "", err
	}

	//response, err :=
	//println("the ping request returns", response)
	return libhttp.Do(req)

}
