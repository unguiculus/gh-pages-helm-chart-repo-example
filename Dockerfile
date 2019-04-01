FROM golang:1.12.1-alpine3.9

WORKDIR /app
COPY main.go .
RUN go build -o hello-helm main.go
ENTRYPOINT ["/app/hello-helm"]
