FROM alpine:3.10
WORKDIR /app
COPY /build/linux/index-api /app/bin/index-api
COPY /usr/local/go/lib/time/zoneinfo.zip /opt/zoneinfo.zip

## 解决alipay获取时区错误
ENV ZONEINFO /opt/zoneinfo.zip
VOLUME [ "/app/config", "/app/logs" ]

ENTRYPOINT ["./bin/index-api"]