# Jaeger Go Instrumentation Example
Five simple Go microservices where the services expose a `/ping` and `/pong` endpoint, instrumented with Jaeger+OpenTracing.

# Getting Started

## Start the example

Starts up the Jaeger all-in-one container, along with our example microservices.
```
$ make start
```

## Run Example-1

Hit `service-a`'s endpoint to trigger the trace.
```
$ curl -w '\n' http://localhost:8081/ping
```

## Validate

Should see `service-a -> service-b -> service-c` on STDOUT.

## Run Example-2

Hit `service-a`'s endpoint to trigger the trace.
```
$ curl -w '\n' http://localhost:8081/pong
```

## Validate

Should see `service-a -> service-d` on STDOUT.

## Run Example-3

Hit `service-d`'s endpoint to trigger the trace.
```
$ curl -w '\n' http://localhost:8084/ping
```

## Validate

Should see `service-d -> service-e` on STDOUT.

#### Go to http://localhost:16686/ and select `service-a` from the "Service" dropdown and click the "Find Traces" button.  

## Stop the example

Stop and remove containers.

```
$ make stop
```
