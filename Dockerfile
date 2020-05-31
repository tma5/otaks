FROM golang:latest as builder

WORKDIR /go/src/github.com/tma5/otaks
COPY . .
RUN cd /go/src/github.com/tma5/otaks && \
    CGO_ENABLED=0 GOOS=linux \
    make compile

##
FROM alpine:latest

RUN apk update && apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder go/src/github.com/tma5/otaks/bin/otaks .
COPY etc /etc/

EXPOSE 8080
EXPOSE 8087

ENTRYPOINT [ "./otaks" ]
CMD [ "serve"]