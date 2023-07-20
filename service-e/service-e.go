package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"log"
	"net/http"
	"ping/lib/tracing"
)

const thisServiceName = "service-e"

// when a ping comes it returns its service name
func main() {

	//initialised tracer for service-b
	// <<<<<<<<<
	tracer, closer := tracing.Init(thisServiceName)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
	// >>>>>>>>>

	println("i am service e")

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {

		// span for this request --the global tracer has all info
		// <<<<<<<<<
		span := tracing.StartPingSpanFromRequest(tracer, r)
		defer span.Finish()
		// >>>>>>>>>

		w.Write([]byte(fmt.Sprintf("%s", thisServiceName)))
	})

	log.Printf("Listening on localhost:8085")
	log.Fatal(http.ListenAndServe(":8085", nil))

}
