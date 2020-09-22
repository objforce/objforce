FROM alpine:3.10
WORKDIR /app
COPY build/linux/meta-srv /app/meta-srv
## 解决alipay获取时区错误
ENV ZONEINFO /opt/zoneinfo.zip

ENTRYPOINT ["/app/meta-srv"]