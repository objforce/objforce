FROM alpine:3.10
WORKDIR /app
COPY /build/linux/index-api /app/bin/index-api

## 解决alipay获取时区错误
ENV ZONEINFO /opt/zoneinfo.zip

ENTRYPOINT ["./bin/index-api"]