FROM golang:alpine AS builder

WORKDIR /src
COPY . /src

RUN go env -w GOPROXY=https://goproxy.cn,direct &&  \
    go build -o wxedge_exporter main.go

FROM alpine

WORKDIR /app
COPY --from=builder /src/wxedge_exporter /app/wxedge_exporter
COPY --from=builder /src/config.yaml /app/config.yaml
EXPOSE 9001

CMD ["/app/wxedge_exporter"]
