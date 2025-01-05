FROM golang:1.23.3-alpine AS builder

ADD . /go/src/helm-chart-playground
WORKDIR /go/src/helm-chart-playground

RUN go mod download
RUN go build -o main main.go

FROM alpine:3

WORKDIR /root

COPY --from=builder /go/src/helm-chart-playground/main ./

ENTRYPOINT ["/root/main"]
