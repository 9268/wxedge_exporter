FROM golang:alpine AS builder

WORKDIR /src
COPY . /src

RUN go env -w GOPROXY=https://goproxy.cn,direct &&  \
    go build -o wxedga_exporter main.go

FROM alpine

WORKDIR /app
COPY --from=builder /src/miwifi_exporter /app/wxedga_exporter
COPY --from=builder /src/config.yaml /app/config.yaml
EXPOSE 9001

CMD ["/app/wxedga_exporter"]
