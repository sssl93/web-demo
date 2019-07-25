FROM busybox

COPY bin/web-demo /bin/
COPY config.yaml /opt/

WORKDIR /opt/

ENTRYPOINT ["web-demo"]