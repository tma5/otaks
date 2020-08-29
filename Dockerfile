FROM golang:latest as builder

WORKDIR /go/src/github.com/tma5/otaks
COPY . .
RUN cd /go/src/github.com/tma5/otaks && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    make compile

##
FROM gcr.io/distroless/base

WORKDIR /app
COPY --from=builder go/src/github.com/tma5/otaks/bin/otaks .
COPY etc /etc/

EXPOSE 8080
EXPOSE 8089

CMD [ "/app/otaks", "--config", "/etc/otaks/otaks.toml", "serve"]