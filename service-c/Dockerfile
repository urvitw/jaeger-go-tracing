ARG port=8083
ARG app=service-c

FROM golang as builder
ARG app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ${app}/ ${app}
COPY lib/ lib

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/${app} ./${app}

# final stage
FROM scratch
ARG app
ARG port
COPY --from=builder /out/${app} /app/
#runs the go code

EXPOSE ${port}
ENTRYPOINT ["/app/service-c"]
