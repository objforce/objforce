FROM alpine:3.10

WORKDIR /app
COPY build/linux/data-srv /app/bin/data-srv
## 解决alipay获取时区错误
ENV ZONEINFO /opt/zoneinfo.zip

ENTRYPOINT ["./bin/data-srv"]
