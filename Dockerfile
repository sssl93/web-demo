FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache curl
COPY bin/web-demo /bin/
COPY config.yaml /opt/

WORKDIR /opt/

ENTRYPOINT ["web-demo"]