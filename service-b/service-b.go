package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"log"
	"net/http"
	"ping/lib/ping"
	"ping/lib/tracing"
)

const thisServiceName = "service-b"

// when a ping comes it returns its service name
func main() {

	//initialised tracer for service-b
	// <<<<<<<<<
	tracer, closer := tracing.Init(thisServiceName)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
	// >>>>>>>>>

	println("i am service b")
	outboundHostPortPing := "service-c:8083"
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {

		// span for this request --the global tracer has all info
		// <<<<<<<<<
		span := tracing.StartPingSpanFromRequest(tracer, r)
		defer span.Finish()
		// >>>>>>>>>

		ctx := opentracing.ContextWithSpan(context.Background(), span)

		response, err := ping.Ping(ctx, outboundHostPortPing)
		if err != nil {
			log.Fatalf("Error occurred: %s", err)
		}
		w.Write([]byte(fmt.Sprintf("%s -> %s", thisServiceName, response)))
	})

	log.Printf("Listening on localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", nil))

}
