package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"log"
	"net/http"
	"ping/lib/ping"
	"ping/lib/pong"
	"ping/lib/tracing"
)

const thisServiceName = "service-a"

// when it is pinged it pings service-b or 8082 port
func main() {

	//initialised tracer for service a
	// <<<<<<<<<
	tracer, closer := tracing.Init(thisServiceName)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
	// >>>>>>>>>

	println("i am service a")

	outboundHostPortPing := "service-b:8082"
	outboundHostPortPong := "service-d:8084"
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

		// start first span
		// <<<<<<<<<
		span := tracing.StartPongSpanFromRequest(tracer, r) // root span
		defer span.Finish()

		//When making a downstream call to service-b, a new context instance is created using the ContextWithSpan function
		//which copies the “server” span’s details into the new context.
		//This is very important, as it ensures the continued lineage of parent-child relationships
		ctx := opentracing.ContextWithSpan(context.Background(), span)

		response, err := pong.Pong(ctx, outboundHostPortPong)
		// >>>>>>>>>
		if err != nil {
			log.Fatalf("Error occurred: %s", err)
		}
		w.Write([]byte(fmt.Sprintf("%s -> %s", thisServiceName, response)))

	})

	log.Printf("Listening on localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
