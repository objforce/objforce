FROM autokit/micro-builder:latest as builder

WORKDIR /opt

RUN go env -w GOPROXY=https://goproxy.cn,direct && go env -w GO111MODULE=on
ADD go.mod .
RUN go mod download

COPY . /opt

RUN make build

FROM alpine:3.10
WORKDIR /app
COPY --from=builder /opt/data-srv /app/bin/data-srv
COPY --from=builder /opt/alipay.crt /etc/ssl/certs/
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /opt/zoneinfo.zip
## 解决alipay获取时区错误
ENV ZONEINFO /opt/zoneinfo.zip

ENTRYPOINT ["./bin/data-srv"]
