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

const thisServiceName = "service-d"

// when it is pinged it pings service-b or 8082 port
func main() {

	//initialised tracer for service d
	// <<<<<<<<<
	tracer, closer := tracing.Init(thisServiceName)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
	// >>>>>>>>>

	println("i am service d")

	outboundHostPortPing := "service-e:8085"
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {

		// start first span
		// <<<<<<<<<
		span := tracing.StartPingSpanFromRequest(tracer, r) // root span
		defer span.Finish()

		//When making a downstream call to service-b, a new context instance is created using the ContextWithSpan function
		//which copies the “server” span’s details into the new context.
		//This is very important, as it ensures the continued lineage of parent-child relationships
		ctx := opentracing.ContextWithSpan(context.Background(), span)

		response, err := ping.Ping(ctx, outboundHostPortPing)
		// >>>>>>>>>
		if err != nil {
			log.Fatalf("Error occurred: %s", err)
		}
		w.Write([]byte(fmt.Sprintf("%s -> %s", thisServiceName, response)))

	})

	http.HandleFunc("/pong", func(w http.ResponseWriter, r *http.Request) {

		// span for this request --the global tracer has all info
		// <<<<<<<<<
		span := tracing.StartPongSpanFromRequest(tracer, r)
		defer span.Finish()
		// >>>>>>>>>

		w.Write([]byte(fmt.Sprintf("%s", thisServiceName)))
	})
	log.Printf("Listening on localhost:8084")
	log.Fatal(http.ListenAndServe(":8084", nil))
}
