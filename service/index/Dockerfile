FROM alpine:3.10
WORKDIR /app
COPY build/linux/index-srv /app/index-srv

## 解决alipay获取时区错误
ENV ZONEINFO /opt/zoneinfo.zip

ENTRYPOINT ["./index-srv"]
