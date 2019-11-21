FROM alpine:3.10

ADD conf /conf
ADD bin/linux_amd64/micro /api-gateway

WORKDIR /
ENTRYPOINT [ "/api-gateway" ]
